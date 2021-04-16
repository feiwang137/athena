/*
@Description:
@Author: fei.wang
@Date: 2021/03/29
*/

package parse_prometheus_rule

import (
	"encoding/json"
	//"io/ioutil"
	"log"
	"net/http"
	"github.com/feiwang137/athena/pkg/models"
)

func ParseRulesFromPrometheus(url string) error  {

	req ,err := http.NewRequest("GET",url,nil)
	if err !=nil{
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err !=nil {
		log.Fatal(err)
	}

	var prp apiFuncResult
	json.NewDecoder(resp.Body).Decode(&prp)

	// 判断数据库是否已经写入数据，如果已经有数据的话，不能重复写入的
	data , err := models.Read()
	if err !=nil{
		return err
	}

	if len(data) > 0{
		log.Println("athena.db have rules data, skip import prometheus rule config ")
		return nil
	}

	err = Rule2DB(&prp)
	if err != nil{
		return err
	}

	err = GenPromRuleConfig()
	if err != nil{
		return err
	}


	return  nil
}