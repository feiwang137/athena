/*
@Description: xxxx
@Author: fei.wang
@Date: 2020/12/09
*/
package server

import (
	"encoding/json"
	"fmt"
	"github.com/feiwang137/athena/pkg/models"
	"github.com/feiwang137/athena/pkg/parse_prometheus_rule"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type Rules []models.MyRule


// 查询rule, 需要传入查询关键字
func QueryRules(c *gin.Context){
	query_range := c.DefaultQuery("query_range","all")

	if query_range == "all"{
		rules, err := models.Read()
		if err != nil{
			log.Fatal(err)
		}

		c.JSON(200,gin.H{
			"data": rules,
		})
	} else {
		rules , err := models.FRead(query_range)
		if err != nil{
			log.Fatal(err)
		}

		c.JSON(200,gin.H{
			"state": "ok",
			"data": rules,
		})
	}
}

// 更新rule，传入一个json
func UpdateRule(c *gin.Context){
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	Rule := make(map[string]interface{})
	c.BindJSON(&Rule)

	if _,ok := Rule["labels"];ok{
		Rule["labels"] = parse_prometheus_rule.StructToString(Rule["labels"])
	}

	if _,ok := Rule["annotations"];ok{
		Rule["annotations"] = parse_prometheus_rule.StructToString(Rule["annotations"])
	}


	// ToDo: to check rules and generate rules.yaml

	err := models.Update(uint(id), &Rule)
	if err !=nil{
		panic(err)
	}
	//
	err = GeneratePromRuleFileAndReloadProm()
	if err !=nil{
		log.Fatalln("Error: ", err)
	}

	c.JSON(200, gin.H{
		"state":"ok",
	})
}

// delete rules
func DeleteRule(c *gin.Context){
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	// ToDo: generate rules.yaml

	err := models.Delete(uint(id))
	if err != nil{
		panic(err)
	}

	err = GeneratePromRuleFileAndReloadProm()
	if err !=nil{
		log.Fatalln("Error: ", err)
	}

	c.JSON(200, gin.H{
		"state":"ok",
	})
}

func CreateRules(c *gin.Context){
	JsonTempRules := make([]map[string]interface{},0)
	c.BindJSON(&JsonTempRules)


	fmt.Printf("labels, T: %T, V: %v \n", JsonTempRules[0]["labels"],JsonTempRules[0]["labels"],)

	data, err := json.MarshalIndent(JsonTempRules,""," ")
	if err != nil{
		log.Println("ERROR: ",err)
		return
	}

	var TempRules models.TempMyRules
	err = json.Unmarshal(data, &TempRules)
	if err != nil{
		log.Println("ERROR: ",err)
		return
	}

	var rules models.MyRules
	for _, temp_rule := range TempRules{
		rule := models.MyRule{
			GroupName: temp_rule.GroupName,
			RuleName: temp_rule.RuleName,
			Type: temp_rule.Type,
			Query: temp_rule.Query,
			Interval: temp_rule.Interval,
			Duration: temp_rule.Duration,
			Labels: parse_prometheus_rule.StructToString(temp_rule.Labels),
			Annotations: parse_prometheus_rule.StructToString(temp_rule.Annotations),
		}
		rules = append(rules, rule)
	}


	//for _, rule := range TempRules{
	//	err := CheckRuleFilesValid(rule)
	//	if err != nil{
	//		log.Println("ERROR: ", err)
	//		// 返回结果，告知错误原因！
	//		c.JSON(200, gin.H{
	//			"state":"error",
	//			"reason": err,
	//		})
	//
	//		return
	//	}
	//
	//}

	// ToDo: to check rules and generate rules.yaml

	err = rules.Create()
	if err != nil{
		panic(err)
	}

	err = GeneratePromRuleFileAndReloadProm()
	if err !=nil{
		log.Fatalln("Error: ", err)
	}

	c.JSON(200, gin.H{
		"state":"ok",
	})
}

func GeneratePromRuleFileAndReloadProm() error{
	err := parse_prometheus_rule.GenPromRuleConfig()
	if  err != nil{
		log.Fatalln("Error generate rule yaml file.",err)
		return err
	}

	err = ReloadPrometheusServer()
	if  err != nil{
		log.Fatalln("Error reload prometheus sever",err)
		return err
	}
	return nil

}