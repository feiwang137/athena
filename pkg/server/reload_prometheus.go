/*
@Description: xxxx
@Author: fei.wang
@Date: 2020/12/09
*/
package server

import (
	"github.com/feiwang137/athena/pkg/utils"
	"io/ioutil"
	"net/http"
	"log"
	//"fmt"
)

//func ReloadPrometheusServer(promServer string) error{
func ReloadPrometheusServer() error{

	serverConfig, err := utils.LoadServerConfig()
	if err != nil{
		panic(err)
	}
	promServer := serverConfig.PrometheusServer

	url := promServer + "/-/reload"
	client := &http.Client{}
	req ,err := http.NewRequest("POST",url,nil)
	if err != nil {
		log.Fatalln("Error ",err)
		return err
	}

	res , err := client.Do(req)
	if err != nil {
		log.Fatalln("Error ",err)
		return err
	}

	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Error ",err)
		return err
	}
	log.Println("reload prometheus server success.")
	return nil
}
