package controller

import (
	. "dungou.cn/datasource"
	. "dungou.cn/util"
	"github.com/astaxie/beego"
	"strings"
)

type PupController struct {
	BaseController
}

var orm Orm
var mssql = Mssql{}
func init() {
	beego.SetLogger("console", "")
	orm.Init()
}

func Insert() {
	list := orm.GetIdList()
	change()
	for _, v := range list {
		name := ToString(v["id"])
		q := ToString(v["jacks"])
		t := ToString(v["pressures"])
		if name == "" {
			continue
		}
		str := "select top 3 * from " + name + " order by AutoKey desc;"
		a := mssql.Query(str)
		if len(a) == 0 {
			continue
		} else {
			for _, v := range a {
				name = getname(name)
				daopan(v,name)
				jingbao(v,name)
				jiaojie(v,name)
				luoxuanji(v,name)
				juejin(v,name)
				tuya(v,name,q,t)
			}
		}
	}
}

func daopan(v map[string]interface{},name string) {
	daopan := Daopan{}
	daopan.Dungou = name
	time,shike:=dateParse(ToString(v["时间"]))
	daopan.Jilutime = time
	daopan.Shike = shike
	daopan.No1 = ToString(v["A20"])
	daopan.No2 = ToString(v["A21"])
	daopan.No3 = ToString(v["A22"])
	daopan.No4 = ToString(v["A23"])
	daopan.Zongniuli = ToString(v["A24"])
	daopan.Waizhou = ToString(v["A23"])
	daopan.Neizhou = ToString(v["A23"])
	daopan.Zuozhuan = ToString(v["A23"])
	daopan.Youzhuan = ToString(v["A23"])
	daopan.Chaowali = ToString(v["A25"])
	daopan.Huixuansudu = ToString(v["A31"])
	Db.Create(&daopan)
}

func jingbao(v map[string]interface{},name string) {
	jingbao := Jingbao{}
	jingbao.Dungou = name
	time,shike:=dateParse(ToString(v["时间"]))
	jingbao.Jilutime = time
	jingbao.Shike = shike
	jingbao.Dyfx  = ToString(v["D43"])
	jingbao.Dpcfh = ToString(v["D44"])
	jingbao.Pzjhz = ToString(v["D45"])
	jingbao.No1lx = ToString(v["D46"])
	jingbao.No2lx = ToString(v["D47"])
	jingbao.PzjJPU= ToString(v["D48"])
	jingbao.Zy    = ToString(v["D49"])
	jingbao.Sb    = ToString(v["D50"])
	jingbao.Jn    = ToString(v["D52"])
	jingbao.No1zz = ToString(v["D53"])
	jingbao.No2zz = ToString(v["D54"])
	jingbao.Yl    = ToString(v["D56"])
	jingbao.Xh    = ToString(v["D57"])
	jingbao.Pdcfh = ToString(v["D58"])
	jingbao.Zzdy  = ToString(v["D59"])
	jingbao.Pddy  = ToString(v["D60"])
	jingbao.Dpdy  = ToString(v["D61"])
	jingbao.Zzhl  = ToString(v["D64"])
	jingbao.Yxhl  = ToString(v["D66"])
	jingbao.Dpcnl = ToString(v["D67"])
	jingbao.Hzj   = ToString(v["D68"])
	jingbao.Mf    = ToString(v["D69"])
	jingbao.Dpss  = ToString(v["D70"])
	Db.Create(&jingbao)
}

func jiaojie(v map[string]interface{},name string) {
	jiaojie := Jiaojie{}
	jiaojie.Dungou = name
	time,shike:=dateParse(ToString(v["时间"]))
	jiaojie.Jilutime = time
	jiaojie.Shike = shike
	jiaojie.Xc1 = ToString(v["A64"])
	jiaojie.Xc2 = ToString(v["A65"])
	jiaojie.Xc3 = ToString(v["A66"])
	jiaojie.Xc4 = ToString(v["A67"])
	jiaojie.Yl1 = ToString(v["A68"])
	jiaojie.Yl2 = ToString(v["A69"])
	jiaojie.Yl3 = ToString(v["A70"])
	jiaojie.Yl4 = ToString(v["A71"])
	jiaojie.Jdsx= ToString(v["A72"])
	jiaojie.Jdzy= ToString(v["A73"])
	Db.Create(&jiaojie)
}

func luoxuanji(v map[string]interface{},name string) {
	luoxuanji := Luoxuanji{}
	luoxuanji.Dungou = name
	time,shike:=dateParse(ToString(v["时间"]))
	luoxuanji.Jilutime = time
	luoxuanji.Shike = shike
	luoxuanji.Hzm = ToString(v["A7"])
	luoxuanji.Yl  = ToString(v["A19"])
	luoxuanji.Zt  = ToString(v["A36"])
	luoxuanji.Sd  = ToString(v["A8"])
	Db.Create(&luoxuanji)
}

func juejin(v map[string]interface{},name string) {
	juejin := Juejin{}
	juejin.Dungou = name
	time,shike:=dateParse(ToString(v["时间"]))
	juejin.Jilutime = time
	juejin.Shike = shike
	juejin.Fy     = ToString(v["A5"])
	juejin.Hz     = ToString(v["A6"])
	juejin.Spq    = ToString(v["A40"])
	juejin.CZq    = ToString(v["A41"])
	juejin.Sph    = ToString(v["A42"])
	juejin.Czh    = ToString(v["A43"])
	juejin.Fw     = ToString(v["A44"])
	juejin.Zjdqh  = ToString(v["A49"])
	juejin.Pmhhy  = ToString(v["A74"])
	juejin.Jhw    = ToString(v["A75"])
	juejin.Dp     = ToString(v["A77"])
	juejin.HBW    = ToString(v["A78"])
	juejin.Ep2    = ToString(v["A79"])
	juejin.Zjzs   = ToString(v["A98"])
	juejin.Zjys   = ToString(v["A99"])
	juejin.Zjyx   = ToString(v["A100"])
	juejin.Zjzx   = ToString(v["A101"])
	juejin.Jjjxz  = ToString(v["D37"])
	juejin.Jjms   = ToString(v["D38"])
	juejin.Zzms   = ToString(v["D39"])
	Db.Create(&juejin)
}

func tuya(v map[string]interface{},name string,q string ,t string) {
	tuya := Tuya{}
	tuya.Dungou = name
	time,shike:=dateParse(ToString(v["时间"]))
	tuya.Jilutime = time
	tuya.Shike = shike
	tuya.Qjdque = q
	tuya.Tuyaque =t
	tuya.Tls    = ToString(v["A26"])
	tuya.Xcs    = ToString(v["A9"])
	tuya.Sds    = ToString(v["A13"])
	tuya.Tlx    = ToString(v["A28"])
	tuya.Xcx    = ToString(v["A11"])
	tuya.Sdx    = ToString(v["A15"])
	tuya.Tlz    = ToString(v["A29"])
	tuya.Xcz    = ToString(v["A12"])
	tuya.Sdz    = ToString(v["A16"])
	tuya.Tly    = ToString(v["A27"])
	tuya.Xcy    = ToString(v["A10"])
	tuya.Sdy    = ToString(v["A14"])
	tuya.Tuya1  = ToString(v["A1"])
	tuya.Tuya2  = ToString(v["A2"])
	tuya.Tuya3  = ToString(v["A3"])
	tuya.Tuya4  = ToString(v["A4"])
	tuya.Tuya5  = ""
	tuya.Ztl    = ToString(v["A30"])
	Db.Create(&tuya)

}


func dateParse(date string)(string,string) {
	a := strings.Split(date, " ")
	return a[0],a[1]
}
func getname(name string)string{
	na := strings.Replace(name, "[dtdg].[dbo].[", "", -1)
	return strings.Replace(na, "]", "", -1)
}

func change() {
	Db.Delete(&Jingbao{})
	Db.Delete(&Jiaojie{})
	Db.Delete(&Daopan{})
	Db.Delete(&Juejin{})
	Db.Delete(&Luoxuanji{})
	Db.Delete(&Tuya{})
}