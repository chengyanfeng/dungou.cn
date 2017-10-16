package controller

import (
	. "dungou.cn/datasource"
	"fmt"
)
type RealController struct {
	BaseController
}

func (this *RealController) Getdaopan() {
	dungou := this.GetString("dungou")
	daopan := []Daopan{}
	Db.Where("dungou = ? and batch = ?", dungou,1).Find(&daopan)
	if len(daopan) == 0  {
		fmt.Println(111111111)
		Db.Where("dungou = ? and batch = ?", dungou,2).Find(&daopan)
	}
	this.EchoJsonMsg(daopan)
}

func (this *RealController) Getjiaojie() {
	dungou := this.GetString("dungou")
	jiaojie := []Jiaojie{}
	Db.Where("dungou = ? and batch = ?", dungou,1).Find(&jiaojie)
	if len(jiaojie) == 0  {
		Db.Where("dungou = ? and batch = ?", dungou,2).Find(&jiaojie)
	}
	this.EchoJsonMsg(jiaojie)
}

func (this *RealController) Getjingbao() {
	dungou := this.GetString("dungou")
	jingbao := []Jingbao{}
	Db.Where("dungou = ? and batch = ?", dungou,1).Find(&jingbao)
	if len(jingbao) == 0 {
		Db.Where("dungou = ? and batch = ?", dungou,2).Find(&jingbao)
	}
	this.EchoJsonMsg(jingbao)
}

func (this *RealController) Getjuejin() {
	dungou := this.GetString("dungou")
	juejin := []Juejin{}
	Db.Where("dungou = ? and batch = ?", dungou,1).Find(&juejin)
	if len(juejin) == 0 {
		Db.Where("dungou = ? and batch = ?", dungou,2).Find(&juejin)
	}
	this.EchoJsonMsg(juejin)
}

func (this *RealController) Getluoxuanji() {
	dungou := this.GetString("dungou")
	luoxuanji := []Luoxuanji{}
	Db.Where("dungou = ? and batch = ?", dungou,1).Find(&luoxuanji)
	if len(luoxuanji) == 0 {
		Db.Where("dungou = ? and batch = ?", dungou,2).Find(&luoxuanji)
	}
	this.EchoJsonMsg(luoxuanji)
}

func (this *ApiController) Gettuya() {
	dungou := this.GetString("dungou")
	tuya := []Tuya{}
	Db.Where("dungou = ? and batch = ?", dungou,1).Find(&tuya)
	if len(tuya) == 0 {
		Db.Where("dungou = ? and batch = ?", dungou,2).Find(&tuya)
	}
	this.EchoJsonMsg(tuya)
}