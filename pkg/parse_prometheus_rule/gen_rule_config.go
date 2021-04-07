/*
@Description: xxxx
@Author: fei.wang
@Date: 2020/12/09
*/

package parse_prometheus_rule

import (
	"log"
	//"github.com/google/martian/log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"strconv"
)

type NRuleGroups struct {
	Groups []NGroup `yaml:"groups"`
}

type NGroup struct {
	Name      string         `yaml:"name"`
	Interval  string         `yaml:"interval"`
	Rules     []interface{}        `yaml:"rules"`
}

// NewAlertRule
type NARule struct {
	Name      	string					`yaml:"alert,omitempty"`
	Expr        string         			`yaml:"expr"`
	For         string         			`yaml:"for,omitempty"`
	Labels      map[string]string		`yaml:"labels"`
	Annotations map[string]string		`yaml:"annotations"`
}

// NewRecordRule
type NRRule struct {
	Name      	string					`yaml:"record,omitempty"`
	Expr        string         			`yaml:"expr"`
	Labels      map[string]string		`yaml:"labels"`
	Annotations map[string]string		`yaml:"annotations,omitempty"`
}

func int2str(time int) (string){
	return strconv.Itoa(time) + "s"
}

func GenPromRuleConfig(pc *apiFuncResult) error{
	rule_data := pc.Data
	rule_groups := rule_data.Groups
	NG := make([]NGroup,len(rule_groups),len(rule_groups))

	for i, rule_group := range rule_groups{

		newGroup := NGroup{
			Name: rule_group.Name,
			Interval: int2str(rule_group.Interval),
		}

		for _, rule := range rule_group.Rules{

			var new_alert_rule NARule
			var new_record_rule NRRule

			if rule.Type == "alerting"{
				new_alert_rule = NARule{
					Name: rule.Alert,
					Expr: rule.Expr,
					For: int2str(rule.For),
					Labels: rule.Labels,
					Annotations: rule.Annotations,
				}
				newGroup.Rules = append(newGroup.Rules, new_alert_rule)

			} else if rule.Type == "recording" {
				new_record_rule = NRRule{
					Name: rule.Alert,
					Expr: rule.Expr,
					Labels: rule.Labels,
					Annotations: rule.Annotations,
				}
				newGroup.Rules = append(newGroup.Rules, new_record_rule)
			}
		}
		NG[i] = newGroup

	}

	NGS := NRuleGroups{Groups: NG}
	data, err := yaml.Marshal(&NGS)
	if err != err{
		return err
	}

	err = ioutil.WriteFile("rules.yaml",data,0644)
	if err != nil{
		log.Printf("save rule to yaml fail, error:%v \n",err)
		return err
	}
	log.Printf("save prometheus config success, rules.yaml\n")
	return nil

}