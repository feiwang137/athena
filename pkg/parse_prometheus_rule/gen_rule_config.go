/*
@Description: xxxx
@Author: fei.wang
@Date: 2020/12/09
*/
package parse_prometheus_rule

import (
	"encoding/json"
	"github.com/feiwang137/athena/pkg/models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strconv"
)

type PromRules struct {
	Groups []Group `yaml:"groups"`
}

type Group struct {
	Name     string        `yaml:"name"`
	Interval string        `yaml:"interval"`
	Rules    []interface{} `yaml:"rules"`
}

// NewAlertRule
type AlertRule struct {
	Name        string            `yaml:"alert"`
	Expr        string            `yaml:"expr"`
	For         string            `yaml:"for"`
	Labels      map[string]string `yaml:"labels"`
	Annotations map[string]string `yaml:"annotations"`
}

// NewRecordRule
type RecordRule struct {
	Name   string            `yaml:"record,omitempty"`
	Expr   string            `yaml:"expr"`
	Labels map[string]string `yaml:"labels"`
}

func Str2Struct(convertStr string) map[string]string {
	bytes := []byte(convertStr)
	var key map[string]string
	json.Unmarshal(bytes, &key)
	return key
}

func Int2str(time int) string {
	return strconv.Itoa(time) + "s"
}

func GenPromRuleConfig() error {

	var PromRulesConfig PromRules

	data, err := models.SpecifyFiled("group_name")

	if err != nil {
		return err
	}

	for _, group := range data {
		var ruleGroup Group

		ruleGroup.Interval = Int2str(group.Interval)
		ruleGroup.Name = group.GroupName
		//按照group name查询rule
		data, err := models.FindByGroupName(group.GroupName)
		if err != nil {
			log.Println(err)
		}

		for _, rule := range data {
			if rule.Type == "alerting" {
				alertRule := AlertRule{
					Name:        rule.RuleName,
					Expr:        rule.Query,
					For:         Int2str(rule.Duration),
					Labels:      Str2Struct(rule.Labels),
					Annotations: Str2Struct(rule.Annotations),
				}
				ruleGroup.Rules = append(ruleGroup.Rules, alertRule)

			} else if rule.Type == "recording" {
				recordRule := RecordRule{
					Name:   rule.RuleName,
					Expr:   rule.Query,
					Labels: Str2Struct(rule.Labels),
				}
				ruleGroup.Rules = append(ruleGroup.Rules, recordRule)
			}
		}
		PromRulesConfig.Groups = append(PromRulesConfig.Groups, ruleGroup)
	}

	err = CreateYamlFile(PromRulesConfig,"/Users/feiwang/prom-data/rules.yml")

	if err != nil {
		log.Printf("save rule to yaml fail, error:%v \n", err)
		return err
	}
	return nil
}

func CreateYamlFile(rules PromRules,rulePath string) error {
	data, err := yaml.Marshal(&rules)
	if err != nil {
		return err
	}

	ruleFilePath := rulePath
	err = ioutil.WriteFile(ruleFilePath, data, 0644)
	if err != nil {
		log.Printf("save rule to yaml fail, error:%v \n", err)
		return err
	}

	log.Printf("Generate %v success.",rulePath)

	return nil
}
