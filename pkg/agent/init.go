/*
@Description:
@Author: fei.wang
@Date: 2021/03/11
*/

package agent

import (
	"github.com/feiwang137/athena/pkg/parse_prometheus_rule"
)

func InitAthena(prometheus_url string) error{
	prometheusRuleUrl := prometheus_url + "/api/v1/rules"
	return parse_prometheus_rule.ParseRulesFromPrometheus(prometheusRuleUrl)
}