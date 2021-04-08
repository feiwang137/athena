/*
@Description: xxxx
@Author: fei.wang
@Date: 2020/12/09
*/

package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	//"log"
	//"os"
)

type MyRules []MyRule

type MyRule struct {
	gorm.Model
	GroupName   string
	RuleName    string
	Type        string
	Query       string
	Interval    int `gorm:"default:15"`
	Duration    int `gorm:"default:0"`
	Labels      string
	Annotations string
}

var db *gorm.DB

func stepDB() {

	autoMigrate := false
	if _, err := os.Stat("athena.db"); err != nil {
		autoMigrate = true
		log.Println("athena.db doesn't exits, create it.")
	} else {
		log.Println("athena.db is exists, loading it.")
	}

	d, err := gorm.Open(sqlite.Open("athena.db"), &gorm.Config{})
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
	return db.Create(&rl).Error
}

// 单条记录删除
func Delete(id uint) (int64, error) {
	r := db.Unscoped().Delete(&MyRule{}, id)
	return r.RowsAffected, r.Error
}

// 单条记录更新
func Update(id uint, myRule MyRule) error {
	db.Model(&MyRule{}).Where(id).Updates(myRule)
	return nil
}

// 模糊查询
func FRead(key string) (MyRules, error) {
	var myRule []MyRule
	fuzzyStr := "%" + key + "%"
	db.Where("rule_name LIKE ?", fuzzyStr).Or("group_name LIKE ?", fuzzyStr).Find(&myRule)
	return myRule, nil
}

// 获取全部rule
func Read() (MyRules, error) {
	var myRule []MyRule
	db.Find(&myRule)
	return myRule, nil
}

func init() {
	stepDB()
}
