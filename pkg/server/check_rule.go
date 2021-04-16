/*
@Description: xxxx
@Author: fei.wang
@Date: 2020/12/09
*/
package server

import (
	"fmt"
	"github.com/feiwang137/athena/pkg/parse_prometheus_rule"
	"github.com/feiwang137/athena/pkg/models"
	"log"
	"os"
	"os/exec"
)

// generate rule config yaml file
func GenRuleFileTemp(rule models.MyRule) (error){
	var PromRulesConfig parse_prometheus_rule.PromRules
	var ruleGroup parse_prometheus_rule.Group

	fmt.Printf("------\n%v\n----",rule)
	ruleGroup.Name = rule.GroupName
	ruleGroup.Interval = parse_prometheus_rule.Int2str(rule.Interval)

	if rule.Type == "alerting" {
		alertRule := parse_prometheus_rule.AlertRule{
			Name:        rule.RuleName,
			Expr:        rule.Query,
			For:         parse_prometheus_rule.Int2str(rule.Duration),
			Labels:      parse_prometheus_rule.Str2Struct(rule.Labels),
			Annotations: parse_prometheus_rule.Str2Struct(rule.Annotations),
		}
		ruleGroup.Rules = append(ruleGroup.Rules, alertRule)
	} else if rule.Type == "recording" {
		recordRule := parse_prometheus_rule.RecordRule{
			Name:   rule.RuleName,
			Expr:   rule.Query,
			Labels: parse_prometheus_rule.Str2Struct(rule.Labels),
		}
		ruleGroup.Rules = append(ruleGroup.Rules, recordRule)
	}

	PromRulesConfig.Groups = append(PromRulesConfig.Groups, ruleGroup)
	err := parse_prometheus_rule.CreateYamlFile(PromRulesConfig,"/Users/feiwang/prom-data/rule_temp.yaml")

	if err != nil {
		log.Printf("save rule to yaml fail, error:%v \n", err)
		return err
	}

	return nil
}

// 调用promtool检查语法, 告知检查结果，删除临时文件

func CheckRuleFilesValid(rule models.MyRule) error{
	// 检查配置文件
	// 告知结果
	// 清理配置文件
	err := GenRuleFileTemp(rule)
	if err != nil{
		return err
	}

	// 调用promtool 检查语法
	cmd := exec.Command("/usr/local/bin/promtool","check","rules", "/Users/feiwang/prom-data/rule_temp.yaml")
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil{
		log.Printf("check rule:%v not passed\n", rule.RuleName)
		return err
	}

	log.Printf("check rule: %v passed.\n", rule.RuleName)
	return nil
}



