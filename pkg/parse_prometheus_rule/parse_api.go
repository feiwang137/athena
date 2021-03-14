/*
@Description:
@Author: fei.wang
@Date: 2021/03/1
*/

package parse_prometheus_rule

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseRulesFromPrometheus(url string) error  {
	//req ,err := http.NewRequest("GET",url,nil)
	//if err !=nil{
	//	log.Fatal(err)
	//}
	//
	//resp, err := http.DefaultClient.Do(req)
	//if err !=nil{
	//	log.Fatal(err)
	//}

	//fmt.Printf("%v\n",resp.Body)

	//var prp apiFuncResult

	//fmt.Println(json.NewDecoder(resp.Body).Decode(&prp))
	//fmt.Println(json.NewDecoder(resp.Body).Decode())


	resp, err := http.Get(url)

	if err != nil{
		return err
	}

	fmt.Println(resp.Header)
	//fmt.Println(json.NewDecoder(resp.Body).Decode(&prp))
	body,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return  nil
}