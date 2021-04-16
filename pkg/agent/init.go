/*
@Description:
@Author: fei.wang
@Date: 2021/03/11
*/

package agent

import (
	"github.com/feiwang137/athena/pkg/parse_prometheus_rule"
	"log"
	"os"
)

func InitAthena(prometheus_url string) error{

	sqliteDBPath := "/Users/feiwang/go/src/athena/athena.db"
	if _, err := os.Stat(sqliteDBPath); err != nil{
		log.Fatalln("already init, skip!")
		return nil
	}

	prometheusRuleUrl := prometheus_url + "/api/v1/rules"
	return parse_prometheus_rule.ParseRulesFromPrometheus(prometheusRuleUrl)
}