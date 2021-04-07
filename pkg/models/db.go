/*
@Description: xxxx
@Author: fei.wang
@Date: 2020/12/09
*/
package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"log"
	"os"

	//"log"
	//"os"
)

type MyRules []MyRule

type MyRule struct {
	gorm.Model
	GroupName string
	RuleName string
	Type string
	Query string
	For	int
	Duration int
	Labels string
	Annotations string
}

var db *gorm.DB

func stepDB(){

	autoMigrate := false
	if _, err := os.Stat("athena.db"); err != nil{
		autoMigrate = true
		log.Println("athena.db doesn't exits, create it.")
	} else {
		log.Println("athena.db is exists, loading it.")
	}

	d, err := gorm.Open(sqlite.Open("athena.db"), &gorm.Config{})
	if err != nil{
		panic("failed to connect database.")
	}

	if autoMigrate{
		d.AutoMigrate(&MyRule{})
	}

	db = d
}

// CRUD
func (rl *MyRules) Create() error{
	return db.Create(&rl).Error
}

func (rl *MyRules) Update() error{
	return nil
}

func (rl *MyRules) Delete() error{
	return nil
}

func (rl *MyRules) Read() error{
	return nil
}



func init(){
	stepDB()
}


