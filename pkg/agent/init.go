/*
@Description:
@Author: fei.wang
@Date: 2021/03/11
*/

package agent

import (
	"fmt"
	"github.com/feiwang137/athena/pkg/parse_prometheus_rule"
)

func InitAthena(prometheus_url string) error{
//1.检查/读取本地的DB文件，如果没有就创建
//2.从prometheus读取rule配置，解析后存到DB里。
//3.生成rule文件。
	var printStdout string = `
1. load local db file, if the db file doesn't exists, will to create.
2. load rule from prometheus.
3. create rule yaml file.
`
	fmt.Println("msg:", printStdout, prometheus_url)
	prometheusRuleUrl := prometheus_url + "/api/v1/rules"
	//prometheusRuleUrl := prometheus_url

	return parse_prometheus_rule.ParseRulesFromPrometheus(prometheusRuleUrl)

}
