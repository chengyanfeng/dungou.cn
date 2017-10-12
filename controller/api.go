package controller

import (
	"bytes"
	. "dungou.cn/datasource"
	. "dungou.cn/util"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"code.google.com/p/mahonia"
	."dungou.cn/def"

	"time"
)

type ApiController struct {
	BaseController
}
var ENC = mahonia.NewEncoder("utf8_general_ci")
const MAX_UPLOAD int64 = 50 * 1024 * 1024
const OWNCOMPANY string = "上海地铁盾构设备工程有限公司"

var url = "http://www.metroshield.com:7070"
var key = []byte("qzQpyDAGqDDaHiOY")

func (this *ApiController) Getdata() {
	urljaxrs := url + "/jaxrs/dataapi/haltinfo/"
	startdate := this.GetString("startdate")
	enddate := this.GetString("enddate")
	p := P{}
	p["startdate"] = startdate
	p["enddate"] = enddate
	res, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	param := encrypt(string(res))
	fmt.Println(urljaxrs + param)

	resp, err := http.Get(urljaxrs + param)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	msg, _ := base64.StdEncoding.DecodeString(string(body))
	origData, err := AesDecrypt(msg, key)
	if err != nil {
		panic(err)
	}

	jd := *JsonDecode(origData)
	this.EchoJsonMsg(jd)
}

func (this *ApiController) Maps() {
	param := this.FormToP("city", "company", "type", "path", "section", "own", "dungou")
	p := make(map[string]interface{})
	for k, v := range param {
		if v != nil {
			p[k] = v
		}
	}
	p["status"]= "1"
	sets := []Dungouset{}
	Db.Where(p).Find(&sets)
	this.EchoJsonMsg(persent(sets))
}

func persent(sets []Dungouset)[]Dungouset{
	re :=[]Dungouset{}
	for _,set:=range sets{
		dungou :=set.Datano
		ring :=set.Ringnum
		daopan :=Daopan{}
		p := make(map[string]interface{})
		p["dungou"] = dungou
		p["batch"] =1
		Db.Where(p).First(&daopan)
		ringnum :=daopan.Ringnum
		percent := 0.0
		if ringnum>0 {
			percent =float64(ringnum)/float64(ring)*100
		}

		s := fmt.Sprintf("%0.1f", percent)
		set.Persent = string(s)+"%"
		re = append(re,set)
	}
	return re
}
func (this *ApiController) Getcommu() {
	commum := []Commum{}
	p:=[]P{}
	Db.Where("batch = ?", 1).Find(&commum)
	if len(commum) == 0 {
		Db.Where(" batch = ?",2).Find(&commum)
	}
	for _, v := range commum {
		s := P{}
		set := Dungouset{}
		name := v.Dungou
		Db.Where("datano = ?", name).Find(&set)
		s["dungou"] = set.Dungou
		s["section"] = set.Section
		s["company"] = set.Company2
		s["date"] = v.Jilutime
		s["time"] = v.Shike
		s["status"] = "断开"
		p = append(p, s)
	}
	this.EchoJsonMsg(p)
}

func (this *ApiController) Getcompany() {
	sets := []Dungouset{}
	companys := make([]string, 0)
	Db.Where("status = ?", 1).Find(&sets)
	for _, v := range sets {
		company := v.Company1
		companys = append(companys, company)
	}
	companys = RemoveDuplicatesAndEmpty(companys)
	this.EchoJson(companys)
}

func (this *ApiController) Gettype() {
	sets := []Dungouset{}
	typelist := make([]string, 0)
	Db.Where("status = ?", 1).Find(&sets)
	for _, v := range sets {
		types := v.Type
		typelist = append(typelist, types)
	}
	typelist = RemoveDuplicatesAndEmpty(typelist)
	this.EchoJson(typelist)
}

func (this *ApiController) Getpath() {
	sets := []Dungouset{}
	paths := make([]string, 0)
	Db.Where("status = ?", 1).Find(&sets)
	for _, v := range sets {
		path := v.Path
		paths = append(paths, path)
	}
	paths = RemoveDuplicatesAndEmpty(paths)
	this.EchoJson(paths)
}

func (this *ApiController) Getdungou() {
	sets := []Dungouset{}
	dungous := make([]string, 0)
	Db.Where("status = ?", 1).Find(&sets)
	for _, v := range sets {
		dungou := v.Dungou
		dungous = append(dungous, dungou)
	}
	dungous = RemoveDuplicatesAndEmpty(dungous)
	this.EchoJson(dungous)
}

func (this *ApiController) Getseclonlat() {
	dungou := this.GetString("dungou")
	set := Dungouset{}
	Db.Where("dungou = ?", dungou).Find(&set)
	section := set.Section
	sec := []Seclonlat{}
	Db.Where("section = ?", section).Find(&sec)
	this.EchoJson(sec)
}

func (this *ApiController) Getprolonlat() {
	dungou := this.GetString("dungou")
	set := Dungouset{}
	Db.Where("dungou = ?", dungou).Find(&set)
	section := set.Section
	prolonlat := []Prolonlat{}
	Db.Where("section = ?", section).Find(&prolonlat)
	this.EchoJson(prolonlat)
}

func (this *ApiController) Getprofile() {
	dungou := this.GetString("dungou")
	fmt.Println(dungou)
	set := Dungouset{}
	Db.Where("dungou = ?", dungou).First(&set)
	fmt.Println(set)
	section := set.Section
	profile := Profile{}
	Db.Where("section = ?", section).First(&profile)
	this.EchoJson(profile)
}

func (this *ApiController) Getsection() {
	sets := []Dungouset{}
	path := this.GetString("path")
	sections := make([]string, 0)
	if path == "" {
		Db.Where("status = ?", 1).Find(&sets)
	} else {
		Db.Where("status = ? and path = ?", 1, path).Find(&sets)
	}
	for _, v := range sets {
		section := v.Section
		sections = append(sections, section)
	}
	sections = RemoveDuplicatesAndEmpty(sections)
	this.EchoJson(sections)
}
func (this *ApiController) Prosafe() {
	dungou := this.GetString("risk")
	param := make(map[string]interface{})
	if dungou != "" {
		param["dungou"] = dungou
	}
	sets := []Dungouset{}
	Db.Find(&sets)
	p := []P{}
	sets = persent(sets)
	for _, v := range sets {
		m := P{}
		section := v.Section
		risks := []Risk{}
		Db.Where("section = ?", section).Find(&risks)
		m["set"] = v
		m["risk"] = risks
		p = append(p,m)
	}
	this.EchoJsonMsg(p)
}
func (this *ApiController) Getrisk() {
	dungou := this.GetString("dungou")
	param := make(map[string]interface{})
	if dungou != "" {
		param["dungou"] = dungou
	}
	param["status"]="1"
	sets := []Dungouset{}
	Db.Where(param).Find(&sets)
	p := []P{}
	sets = persent(sets)
	for _, v := range sets {
		m := P{}
		section := v.Section
		risks := []Risk{}
		Db.Where("section = ?", section).Find(&risks)
		m["set"] = v
		m["risk"] = risks
		p = append(p,m)
	}
	this.EchoJsonMsg(p)
}

func (this *ApiController) Getsediment() {
	dungou := this.GetString("dungou")
	sediment := []Sediment{}
	if dungou != "" {
		Db.Where("batch = ? and dungou =? ", 1,dungou).Find(&sediment)
		if len(sediment) == 0  {
			Db.Where("batch = ? and dungou =?",2,dungou).Find(&sediment)
		}
	}else{
		Db.Where("batch = ?", 1).Find(&sediment)
		if len(sediment) == 0  {
			Db.Where("batch = ?",2).Find(&sediment)
		}
	}
	this.EchoJsonMsg(sediment)
}

func (this *ApiController) Getdaopan() {
	dungou := this.GetString("dungou")
	daopan := []Daopan{}
	Db.Where("dungou = ? and batch = ?", dungou,1).Find(&daopan)
	if len(daopan) == 0  {
		fmt.Println(111111111)
		Db.Where("dungou = ? and batch = ?", dungou,2).Find(&daopan)
	}
	this.EchoJsonMsg(daopan)
}
func (this *ApiController) Getjiaojie() {
	dungou := this.GetString("dungou")
	jiaojie := []Jiaojie{}
	Db.Where("dungou = ? and batch = ?", dungou,1).Find(&jiaojie)
	if len(jiaojie) == 0  {
		Db.Where("dungou = ? and batch = ?", dungou,2).Find(&jiaojie)
	}
	this.EchoJsonMsg(jiaojie)
}

func (this *ApiController) Getjingbao() {
	dungou := this.GetString("dungou")
	jingbao := []Jingbao{}
	Db.Where("dungou = ? and batch = ?", dungou,1).Find(&jingbao)
	if len(jingbao) == 0 {
		Db.Where("dungou = ? and batch = ?", dungou,2).Find(&jingbao)
	}
	this.EchoJsonMsg(jingbao)
}

func (this *ApiController) Getjuejin() {
	dungou := this.GetString("dungou")
	juejin := []Juejin{}
	Db.Where("dungou = ? and batch = ?", dungou,1).Find(&juejin)
	if len(juejin) == 0 {
		Db.Where("dungou = ? and batch = ?", dungou,2).Find(&juejin)
	}
	this.EchoJsonMsg(juejin)
}

func (this *ApiController) Getluoxuanji() {
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

//登陆
func (this *ApiController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")
	password = Md5(password, Md5Salt)
	user := User{}
	p := make(map[string]string)
	p["username"] = username
	Db.Where("username = ? ", username).First(&user)
	if user.Username == "" {
		this.EchoJsonErr("用户不存在")
		this.StopRun()
	}

	if user.Password != password {
		fmt.Println(user.Password)
		fmt.Println(password)
		this.EchoJsonErr("密码错误")
		this.StopRun()
	}
	S(user.Grade,user.Grade)
	user.Password=""
	this.EchoJsonMsg(user)
}

func (this *ApiController)Exit(){
	grade:=this.GetString("grade")
	log:=Del(grade)
	if log=="ok"{
		this.EchoJsonMsg("ok")
	}else {
		this.EchoJsonErr("error")
	}

	fmt.Println(log)

}

//添加
func (this *ApiController)Adduser(){

	user:=User{}
	userfind:=User{}
	username := this.GetString("username")
	Db.Where("username = ? ", username).First(&user)
	if !IsEmpty(user.Username) {
		this.EchoJsonErr("用户已注册")
	}else {
		password := this.GetString("password")
		password = Md5(password, Md5Salt)
		role:=this.GetString("role")
		companyid:=this.GetString("companyid")
		id:=ToInt(this.GetString("id"))
		user.Id=id
		user.Username=username
		user.Password=password
		user.Role=role
		user.Companyid=companyid
		user.Grade=this.GetString("grade")
		user.Grade= Md5(username+ToString(time.Now()), Md5Salt)
		Db.Create(&user)
		Db.Where("username = ? ", username).First(&userfind)
		if !IsEmpty(userfind.Username) {
			this.EchoJsonMsg("插入成功")
		}else {
			this.EchoJsonErr("插入失败")
		}
	}
}
//修改
func (this *ApiController)Updateuser(){
	user:=User{}
	param:=make(map[string]interface{})
	id:=this.GetString("id")
	p := this.FormToP("password", "role","companyid","username")
	for k,v:=range p{
		if v!=nil{
			param[k]=v
		}
	}
	db:=Db.Model(&user).Where("id = ?", id).Updates(param)
	if strings.Fields(ToString(db))[1]=="<nil>"{
		if strings.Fields(ToString(db))[2]!="0" {
			this.EchoJsonMsg("更新成功")
		}else{
			this.EchoJsonErr("更新失败")
		}

	}else
	{this.EchoJsonErr("更新失败")}



}

//密码修改

func (this *ApiController) Updatepassword(){
	user:=User{}
	grade:=this.GetString("grade")
	passwordnew:=this.GetString("passwordnew")
	passwordold:=this.GetString("passwordold")
	passwordold = Md5(passwordold, Md5Salt)
	passwordnew = Md5(passwordnew, Md5Salt)
	if err:=Db.Where("grade = ?", grade).Find(&user).Error;err!=nil{
		this.EchoJsonErr("修改失败")
	}else{
		if user.Password==passwordold{
		if err:=Db.Model(&user).Where("grade = ?", grade).Update("password", passwordnew).Error; err != nil {
			this.EchoJsonErr("修改失败")

		} else {
			this.EchoJsonMsg("修改成功")
		}
		}else {
			this.EchoJsonMsg("原始密码错误，请重新输入")
		}
	}
}
//查询
func (this *ApiController)Finduser(){
	var db interface{}
	users:=[]User{}
	role := this.GetString("role")
	chilrole := this.GetString("chilrole")
	username:=this.GetString("username")
	p := this.FormToP("username","chilrole")
	param:=make(map[string]interface{})
	for k,v :=range p{
		if v!=nil{
			if k=="chilrole"{
			param["role"]=v
			}else {
				param[k]=v
			}
		}
	}
	fmt.Println(param)
	if role=="2"{
		if IsEmpty(chilrole){
			if IsEmpty(username){
				db=Db.Where("role in (?)", []string{"4", "5"}).Find(&users)
			}else {
		db=Db.Where("role in (?) AND username = ?  ", []string{"4", "5"},username).Find(&users)
		}}else {
			if(chilrole=="4"||chilrole=="5"){
				db=Db.Where(param).Find(&users)
			}else{
				this.EchoJsonErr("没有权限访问")
				this.StopRun()
			}

		}
		}else {
		db=Db.Where(param).Find(&users)
	}


	if strings.Fields(ToString(db))[1]=="<nil>"{
		if strings.Fields(ToString(db))[2]!="0" {
			for k,_:=range users{
				users[k].Password=""
			}
			this.EchoJsonMsg(users)
		}else{
			this.EchoJsonMsg(false)
		}

	}else
	{this.EchoJsonMsg(false)}
}
//删除
func (this * ApiController)Deletuser(){
	id := this.GetString("id")

	db:=Db.Where("id = ?", id).Delete(User{})
	fmt.Println(db)
	if strings.Fields(ToString(db))[2]=="<nil>"{
		if strings.Fields(ToString(db))[3]!="0" {
			this.EchoJsonMsg("删除成功")
		}else{
			this.EchoJsonErr("删除失败")
		}

	}else
	{this.EchoJsonErr("删除失败")}
}

func (this *ApiController) Upload() {
	f, h, err := this.GetFile("bin")
	defer func() {
		if f != nil {
			f.Close()
		}
		if err := recover(); err != nil {
			Error("Upload", err)
		}
	}()

	if err != nil {
		Error("Upload", err)
		this.EchoJsonErr(err.Error())
	} else {
		ext := ToString(Pathinfo(h.Filename)["extension"])
		if !InArray(ext, []string{"png", "jpg", "jpeg", "bmp", "gif", "json", "csv", "xlsx", "xls", "txt", "xml"}) {
			this.EchoJsonErr("文件类型不合法")
		}
		var buff bytes.Buffer
		fileSize, _ := buff.ReadFrom(f)
		if fileSize > MAX_UPLOAD {
			this.EchoJsonErr("文件尺寸不得大于", MAX_UPLOAD)
		} else {
			md5 := Md5(buff.Bytes())
			filename := JoinStr(md5, ".", ext)
			updir := "upload"
			locfile := updir + "/" + filename
			exist := FileExists(locfile)
			if !exist {
				this.SaveToFile("bin", locfile)
			} else {
				Debug("File exists, skip")
			}
			r := P{}
			if ext == "csv" {
				// auto convert gbk to utf-8
				cmd := fmt.Sprintf("enca -L zh_CN -x UTF-8 %v", locfile)
				Exec(cmd)
			}
			switch ext {
			case "xls", "xlsx":
				excel := Excel{}
				sheets, _ := excel.List(locfile)
				files, _ := excel.Xsl2Csv(locfile, JoinStr(updir, "/", filename, ".csv"))
				r["sheets"] = sheets
				r["files"] = files
				if len(files) == 1 {
					r["url"] = files[0]
				}
			default:
				r["url"] = updir + "/" + filename
			}

			r["ext"] = ext
			r["size"] = fileSize
			this.EchoJsonMsg(r)
		}
	}
}

func (this *ApiController) Pub() {
	url := this.GetString("url")
	table := this.GetString("table")
	section := this.GetString("section")
	enc := mahonia.NewEncoder("UTF-8")
	section = enc.ConvertString(section)
	pg := Mysql{}

	if table == "profile" {
		profile := Profile{}
		profile.Section = section
		Db.Where("section = ?", section).Delete(Profile{})
		profile.Url = url
		Db.Create(profile)
	} else if table == "dungouset" {
		file, err := os.Open(url)
		if err != nil {
			this.EchoJsonErr("Error:", err)
			return
		}
		defer file.Close()
		reader := csv.NewReader(file)
		k := 0
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				this.EchoJsonErr(err)
				return
			}
			if k != 0 {
				inserSet(record)
			}
			k++
		}
	} else if table == "rtinfo" || table == "seclonlat" || table == "prolonlat" || table == "risk" {
		_, e := pg.LoadCsv(url, table, ",")
		if e != nil {
			this.EchoJsonErr(e)
			return
		}
	}
	this.EchoJson("200")

}

func encrypt(param string) string {
	result, err := AesEncrypt([]byte(param), key)
	if err != nil {
		panic(err)
	}
	param = base64.StdEncoding.EncodeToString(result)
	param = strings.Replace(param, "/", "-", -1)
	return param
}
//上报信息
func (this *ApiController) Upmessage() {
	message := Message{}
	message.Username = this.GetString("username")
	message.Companyid = this.GetString("line")
	message.Section=this.GetString("section")
	message.Companyid = this.GetString("companyid")
	message.Img = this.GetString("img")
	message.Text = this.GetString("text")
	message.Dungou=this.GetString("dungou")
	message.Ringnum=this.GetString("ringnum")
	message.Type=this.GetString("type")
	message.Schedule=this.GetString("schedule")
	message.Date = this.GetString("date")
	db := Db.Create(&message)
	fmt.Println(db)
	if strings.Fields(ToString(db))[1] == "<nil>" {
		if strings.Fields(ToString(db))[2] != "0" {
			this.EchoJsonMsg("上报成功")
		} else {
			this.EchoJsonErr("上报失败")
		}
	} else {
		this.EchoJsonErr("上报失败")
	}
}
//显示上报信息
func (this *ApiController) Findmessage(){

	messages:=[]Message{}
	p := this.FormToP("username","dungou","type","ringnum")
	param:=make(map[string]interface{})
	for k,v :=range p{
		if v!=nil{
			param[k]=v
		}
	}
	db:=Db.Where(param).Find(&messages)
	fmt.Println(db)

	if strings.Fields(ToString(db))[1]=="<nil>"{
		if strings.Fields(ToString(db))[2]!="0" {

			this.EchoJsonMsg(messages)
		}else{
			this.EchoJsonMsg(messages)
		}

	}else
	{this.EchoJsonErr("查询失败")}
}

func inserSet(record []string) {
	enc := mahonia.NewEncoder("UTF-8")
	set := Dungouset{}
	p := P{}
	dungou := enc.ConvertString(record[3])
	status := "1"
	p["dungou"] = dungou
	p["status"] = status
	Db.Table("dungouset").Where("dungou = ? and status = ?", dungou, status).Updates(P{"status": "0"})
	log.Println(enc.ConvertString(record[0]))
	set.Project = enc.ConvertString(record[0])
	set.Path = enc.ConvertString(record[1])
	set.Section = enc.ConvertString(record[2])
	set.Dungou = enc.ConvertString(record[3])
	set.Type = enc.ConvertString(record[4])
	set.Company1 = enc.ConvertString(record[5])
	set.Company2 = enc.ConvertString(record[6])
	set.Client = enc.ConvertString(record[7])
	set.Datano = enc.ConvertString(record[8])
	set.Pressures =  ToInt(record[9])
	set.Jack =  ToInt(record[10])
	set.Ringnum =  ToInt(record[11])
	set.Lon = enc.ConvertString(record[12])
	set.Lat = enc.ConvertString(record[13])
	set.Schedule = enc.ConvertString(record[15])

	set.Status = status
	if record[14]=="上海" {
		set.City = enc.ConvertString(record[14])
	}else {
		set.City = "全国"
	}
	if enc.ConvertString(record[5]) == OWNCOMPANY {
		set.Own = "1"
	} else {
		set.Own = "0"
	}
	Db.Create(set)
}

//显示备注
func (this *ApiController)Findremark(){
	remark:=[]Remark{}
	p := this.FormToP("username","companyid","messageid","text")
	param:=make(map[string]interface{})
	for k,v :=range p{
		if v!=nil{
			param[k]=v
		}
	}
	db:=Db.Where(param).Find(&remark)
	fmt.Println(db)

	if strings.Fields(ToString(db))[1]=="<nil>"{
		if strings.Fields(ToString(db))[2]!="0" {

			this.EchoJsonMsg(remark)
		}else{
			this.EchoJsonMsg(remark)
		}

	}else
	{this.EchoJsonErr("查询失败")}
}
//添加备注
func (this *ApiController) Upremark() {
	remark := Remark{}
	remark.Username = this.GetString("username")
	remark.Companyid = this.GetString("companyid")
	remark.Messageid, _ = strconv.Atoi((this.GetString("messageid")))
	remark.Text = this.GetString("text")
	remark.Date = this.GetString("date")
	db := Db.Create(&remark)
	if strings.Fields(ToString(db))[1] == "<nil>" {
		if strings.Fields(ToString(db))[2] != "0" {
			this.EchoJsonMsg("添加备注成功")
		} else {
			this.EchoJsonErr("添加备注失败")
		}
	} else {
		this.EchoJsonErr("添加备注失败")
	}

}

