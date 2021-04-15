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

	//fmt.Println(prp.Data)
	// save prometheus config to mysql.
	err = Rule2DB(&prp)
	if err != nil{
		return err
	}

	// save file to yaml.
	//fmt.Println(prp)
	err = GenPromRuleConfig()
	if err != nil{
		return err
	}


	return  nil
}