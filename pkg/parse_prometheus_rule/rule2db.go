/*
@Description: xxxx
@Author: fei.wang
@Date: 2020/12/09
*/
package parse_prometheus_rule

import (
	"encoding/json"
	"github.com/feiwang137/athena/pkg/models"
	"log"
)

func structToString(data map[string]string) string {
	byteData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(byteData)
}

func Rule2DB(prp *apiFuncResult) error {
	// 解析rules,组装成一个slice，然后调用create方法批量的创建

	rData := prp.Data
	var tempRules models.MyRules
	var Rules *models.MyRules
	for _, gr := range rData.Groups {
		for _, r := range gr.Rules {
			rule := models.MyRule{
				GroupName:   gr.Name,
				RuleName:    r.Alert,
				Type:        r.Type,
				Query:       r.Expr,
				For:         r.For,
				Duration:    gr.Interval,
				Labels:      structToString(r.Labels),
				Annotations: structToString(r.Annotations),
			}
			tempRules = append(tempRules, rule)
			Rules = &tempRules
		}
	}

	err := Rules.Create()
	if err != nil {
		return err
	}

	log.Println("save prometheus config to DB（Sqlite） success.")
	return nil

}
