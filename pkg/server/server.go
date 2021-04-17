/*
@Description: xxxx
@Author: fei.wang
@Date: 2020/12/09
*/

package server

// 在这里定义相关的接口
// 所有接口均返回json格式
// 这里启动端口需要指定

// 每次进行CUD操作都需要生成最新的配置文件， 并检查配置文件的合法性 ，然后需要reload prom

import (
	//"log"
	"github.com/feiwang137/athena/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"

	//"net/http"
)

const (
	queryRules = "/query_rules"
	updateRule = "/update_rule/:id"
	deleteRule = "/delete_rule/:id"
	createRule = "/create_rule"
)

func PromServer(){
	serverConfig, err := utils.LoadServerConfig()
	if err != nil{
		log.Println("load Athena server listen adress faild, error: ", err)
		panic(err)
	}

	listenAddress := serverConfig.ListenAddress
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET(queryRules, QueryRules)
		v1.POST(updateRule, UpdateRule)
		v1.DELETE(deleteRule, DeleteRule)
		v1.POST(createRule, CreateRules)
	}
	router.Run(listenAddress)
}

