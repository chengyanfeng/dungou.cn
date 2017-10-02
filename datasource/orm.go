package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "dungou.cn/util"
	"fmt"
	"errors"
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

func (Rtinfo) TableName() string {
	return "rtinfo"
}

func (Profile) TableName() string {
	return "profile"
}

func (Seclonlat) TableName() string {
	return "seclonlat"
}

func (Prolonlat) TableName() string {
	return "prolonlat"
}

func (Commum) TableName() string {
	return "commum"
}

func (Sediment) TableName() string {
	return "sediment"
}

func (Message) TableName() string {
	return "message"
}

func (User)TableName() string  {
	return "user"
}

func (Risk) TableName() string {
	return "risk"
}

func (Remark) TableName() string {
	return "remark"
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
	if Db.HasTable("dungouset") {

	} else {
		Db.CreateTable(&Dungouset{})
	}
	if Db.HasTable("rtinfo") {

	} else {
		Db.CreateTable(&Rtinfo{})
	}
	if Db.HasTable("daopan") {

	} else {
		Db.CreateTable(&Daopan{})
	}
	if Db.HasTable("jingbao") {

	} else {
		Db.CreateTable(&Jingbao{})
	}
	if Db.HasTable("user") {

	} else {
		Db.CreateTable(&User{})
	}
	if Db.HasTable("message") {

	} else {
		Db.CreateTable(&Message{})
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
	if Db.HasTable("profile") {

	} else {
		Db.CreateTable(&Profile{})
	}
	if Db.HasTable("seclonlat") {

	} else {
		Db.CreateTable(&Seclonlat{})
	}
	if Db.HasTable("prolonlat") {

	} else {
		Db.CreateTable(&Prolonlat{})
	}
	if Db.HasTable("commum") {

	} else {
		Db.CreateTable(&Commum{})
	}
	if Db.HasTable("sediment") {

	} else {
		Db.CreateTable(&Sediment{})
	}
	if Db.HasTable("risk") {

	} else {
		Db.CreateTable(&Risk{})
	}

	if Db.HasTable("remark") {

	} else {
		Db.CreateTable(&Remark{})
	}
}

func (this *Orm) GetIdList() []map[string]interface{} {
	set := new([]Dungouset)
	Db.Where("status = ?", 1).Find(&set)
	list := make([]map[string]interface{}, 0)
	for _, v := range *set {
		a :=make(map[string]interface{},0)
		a["id"]=v.Datano
		a["type"]=v.Type
		a["status"]=v.Status
		a["jacks"]=v.Jack
		a["pressures"]=v.Pressures
		list = append(list, a)
	}
	return list
}

type Mysql struct {
	P P
}
func (this *Mysql) RunCmd(csv string,table string) (r string, e error) {
	username := Conn["username"]
	database := Conn["name"]
	tpl := `mysql --local-infile=1 -u %v %v -e "LOAD DATA LOCAL INFILE '%v' replace INTO TABLE %v FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n'  ignore 1 lines "`
	compose := fmt.Sprintf(tpl, username, database, csv,table)
	r, e = Exec(compose)
	return
}

func (this *Mysql) LoadCsv(csv string, table string, split string) (r string, e error) {
	if IsEmpty(csv) || IsEmpty(table) {
		e = errors.New(fmt.Sprintf("Invalid csv %v, table %v, database", csv, table))
		return
	}
	return this.RunCmd(csv,table)
}
