package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"time"
	. "dungou.cn/util"
	"fmt"
)

var Db *gorm.DB

type Orm struct {
}

func (Dungouset) TableName() string {
	return "dungouset"
}

func (Daopan) TableName() string {
	return "daopan"
}

func (Jingbao) TableName() string {
	return "jingbao"
}

func (Jiaojie) TableName() string {
	return "jiaojie"
}

func (Luoxuanji) TableName() string {
	return "luoxuanji"
}

func (Juejin) TableName() string {
	return "juejin"
}

func (Tuya) TableName() string {
	return "tuya"
}

func (this *Orm) Init() {
	var err error
	conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		Conn["username"],
		Conn["password"],
		Conn["host"],
		Conn["port"],
		Conn["name"],
	)
	Debug("db conn", conn)
	Db, err = gorm.Open("mysql", conn)
	if err != nil {
		panic(err.Error())
	}
	if Db.HasTable("daopan") {

	} else {
		Db.CreateTable(&Daopan{})
	}
	if Db.HasTable("jingbao") {

	} else {
		Db.CreateTable(&Jingbao{})
	}
	if Db.HasTable("jiaojie") {

	} else {
		Db.CreateTable(&Jiaojie{})
	}
	if Db.HasTable("luoxuanji") {

	} else {
		Db.CreateTable(&Luoxuanji{})
	}
	if Db.HasTable("juejin") {

	} else {
		Db.CreateTable(&Juejin{})
	}
	if Db.HasTable("tuya") {

	} else {
		Db.CreateTable(&Tuya{})
	}

}

func (this *Orm) GetIdList() []map[string]interface{} {
	set := new([]Dungouset)
	Db.Where("status = ?", 1).Find(&set)
	list := make([]map[string]interface{}, 0)
	for _, v := range *set {
		a :=make(map[string]interface{},0)
		a["id"]=v.Id
		a["type"]=v.Type
		a["status"]=v.Status
		a["nowsta"]=v.Nowsta
		a["jacks"]=v.Jacks
		a["pressures"]=v.Pressures
		list = append(list, a)
	}
	return list
}
