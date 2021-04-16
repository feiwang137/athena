/*
@Description: xxxx
@Author: fei.wang
@Date: 2020/12/09
*/
package server

import (
	"io/ioutil"
	"net/http"
	"log"
	//"fmt"
)

//func ReloadPrometheusServer(promServer string) error{
func ReloadPrometheusServer() error{


	//url := "http://" + promServer + "/-/reload"
	url := "http://" + "0.0.0.0:9090" + "/-/reload"
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
