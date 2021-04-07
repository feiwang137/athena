/*
@Description: xxxx
@Author: fei.wang
@Date: 2020/12/09
*/
package parse_prometheus_rule

import (
	"gopkg.in/yaml.v2"
	"log"
)

var data = `
a: Easy 
B: 
  c: 2
  d: [3,4]
`



type BS struct {
	RenamedC int `yaml:"c"`
	D []int `yaml:",flow"`
}

type T struct {
	A string
	B BS
}

func main()  {

	t := T{}
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil{
		log.Fatalf("error: %v",err)
	}

	
}