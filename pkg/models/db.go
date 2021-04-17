/*
@Description: xxxx
@Author: fei.wang
@Date: 2020/12/09
*/

package models

import (
	//"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"github.com/feiwang137/athena/pkg/utils"
)

type MyRules []MyRule

type MyRule struct {
	gorm.Model
	GroupName   string `json:"group_name"`
	RuleName    string `json:"rule_name"`
	Type        string `json:"type"`
	Query       string `json:"query"`
	Interval    int    `json:"interval",gorm:"default:15"`
	Duration    int    `json:"duration",gorm:"default:0"`
	Labels      string `json:"labels"`
	Annotations string `json:"annotations"`
}


type TempMyRules []TempMyRule
type TempMyRule struct {
	gorm.Model
	GroupName   string `json:"group_name"`
	RuleName    string `json:"rule_name"`
	Type        string `json:"type"`
	Query       string `json:"query"`
	Interval    int    `json:"interval",gorm:"default:15"`
	Duration    int    `json:"duration",gorm:"default:0"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}
var db *gorm.DB

func stepDB() {

	autoMigrate := false

	serverConfig, err := utils.LoadServerConfig()
	if err !=nil{
		panic(err)
	}

	if _, err := os.Stat(serverConfig.SqliteDBPath); err != nil {
		autoMigrate = true
		log.Println("first run, will to init DB.")
	}
	d, err := gorm.Open(sqlite.Open(serverConfig.SqliteDBPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	}

	if autoMigrate {
		d.AutoMigrate(&MyRule{})
	}

	db = d
}

// 支持批量添加和单独记录的添加
func (rl *MyRules) Create() error {
	stepDB()
	return db.Create(&rl).Error
}


func Create(rl interface{}) error {
	stepDB()
	//db.Table("my_rules")
	//return db.Create(&rl).Error
	return db.Model(&MyRule{}).Create(&rl).Error
}

// 单条记录删除
func Delete(id uint) error {
	stepDB()
	r := db.Unscoped().Delete(&MyRule{}, id)
	return r.Error
}

// 单条记录更新
func Update(id uint, myRule interface{}) error {
	stepDB()
	db.Model(&MyRule{}).Where(id).Updates(myRule)
	return nil
}

// 模糊查询
func FRead(key string) (MyRules, error) {
	stepDB()
	var myRules []MyRule
	fuzzyStr := "%" + key + "%"
	db.Where("rule_name LIKE ?", fuzzyStr).Or("group_name LIKE ?", fuzzyStr).Find(&myRules)
	return myRules, nil
}

// 获取全部rule
func Read() (MyRules, error) {
	stepDB()
	var myRules []MyRule
	db.Find(&myRules)
	return myRules, nil
}

func FindByGroupName(groupName string) (MyRules, error) {
	stepDB()
	var myRules []MyRule
	db.Where("group_name = ?", groupName).Find(&myRules)
	return myRules, nil
}

func SpecifyFiled(query string) (MyRules, error) {
	stepDB()
	var myRules []MyRule
	db.Select(query, "interval").Group(query).Find(&myRules)
	return myRules, nil
}

