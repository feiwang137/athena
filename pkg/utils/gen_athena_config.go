/*
@Description: xxxx
@Author: fei.wang
@Date: 2020/12/09
*/

package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ServerConfig struct {
	PrometheusServer   string `yaml:"PrometheusServer"`
	ListenAddress string `yaml:"ServerListenAdress"`
	RulesPath          string `yaml:"RulesPath"`
	PromToolPath       string `yaml:"PromToolPath"`
	SqliteDBPath       string `yaml:"SqliteDBPath"`
}

// 生成配置文件
func GenServerConfig(args ...string) error{

	var serverConfig ServerConfig

	serverConfig.PrometheusServer = args[0]
	serverConfig.SqliteDBPath = args[1]
	serverConfig.RulesPath = args[2]
	serverConfig.ListenAddress = args[3]
	serverConfig.PromToolPath = args[4]

	data , err := yaml.Marshal(&serverConfig)
	if err != nil {
		return err
	}

	athenaConfigPath := args[5]
	err = ioutil.WriteFile(athenaConfigPath, data, 0644)
	if err != nil {
		log.Printf("save rule to yaml fail, error:%v \n", err)
		return err
	}

	log.Printf("Generate %v success.", athenaConfigPath)
	return nil

}

// 解析配置文件
var serverConfig *ServerConfig
func LoadServerConfig() error{
	/*
	1.读取配置文件内容
	2.定义一个var，type是对应的Struct
	3.yaml.Unmarshal([]byte(data), &m)
	*/

	return nil

}

