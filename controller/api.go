package controller

import (
	"bytes"
	. "dungou.cn/datasource"
	."dungou.cn/util"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	//."dungou.cn/def"

)

type ApiController struct {
	BaseController
}

const MAX_UPLOAD int64 = 50 * 1024 * 1024

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

func (this *ApiController) Getdaopan() {
	dungou := this.GetString("dungou")
	daopan := Daopan{}
	Db.Where("dungou = ? ", dungou).First(&daopan)
	this.EchoJsonMsg(daopan)
}
//登陆
func (this *ApiController)Login(){
	username := this.GetString("username")
	password := this.GetString("password")
	/*password = Md5(password, Md5Salt)*/
	user:=User{}
	p:=make(map[string]string)
	p["username"]=username
	Db.Where("username = ? ", username).First(&user)
	if user.Username==""{
		this.EchoJsonErr("用户不存在")
		this.StopRun()
	}
	if user.Password!=password{
		fmt.Println(user.Password)
		fmt.Println(password)
		this.EchoJsonErr("密码错误")
		this.StopRun()
	}
	k,_:=json.Marshal(user)
	this.EchoJsonMsg(JsonDecode([]byte(strings.ToLower(string(k)))))
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
		/*password = Md5(password, Md5Salt)*/
		role:=this.GetString("role")
		companyid:=this.GetString("companyid")
		id:=ToInt(this.GetString("id"))
		user.Id=id
		user.Username=username
		user.Password=password
		user.Role=role
		user.Companyid=companyid
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
	p := this.FormToP("password", "role","companyid","username")
	for k,v:=range p{
		if v!=nil{
			param[k]=v
		}
	}
	db:=Db.Model(&user).Where("username = ?", p["username"]).Updates(param)
	if strings.Fields(ToString(db))[1]=="<nil>"{
		if strings.Fields(ToString(db))[2]!="0" {
			this.EchoJsonMsg("更新成功")
		}else{
			this.EchoJsonErr("更新失败")
		}

	}else
	{this.EchoJsonErr("更新失败")}



}
//查询
func (this *ApiController)Finduser(){
	users:=[]User{}
	p := this.FormToP("username","role","companyid","id")
	param:=make(map[string]interface{})
	for k,v :=range p{
		if v!=nil{
			param[k]=v
		}
	}
	db:=Db.Where(param).Find(&users)
	fmt.Println(db)

	if strings.Fields(ToString(db))[1]=="<nil>"{
		if strings.Fields(ToString(db))[2]!="0" {
			k,_:=json.Marshal(users)
			fmt.Println("fanhuideshuju")
			fmt.Println(JsonDecode([]byte(strings.ToLower(string(k)))))
			this.EchoJsonMsg(strings.ToLower(string(k)))
		}else{
			this.EchoJsonErr("查询失败")
		}

	}else
	{this.EchoJsonErr("查询失败")}
}
//删除
func (this * ApiController)Deletuser(){

	username := this.GetString("username")
	db:=Db.Where("username = ?", username).Delete(User{})
	fmt.Println(db)
	a:=*db.Value.(map[string]interface{})
	fmt.Println(a)
	if strings.Fields(ToString(db))[1]=="<nil>"{
		if strings.Fields(ToString(db))[2]!="0" {
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
			updir := ":"
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
func (this *ApiController) Upmessage(){
	message:=Message{}
	message.Username=this.GetString("username")
	message.Companyid=this.GetString("companyid")
	message.Img=this.GetString("img")
	message.Text=this.GetString("text")
	message.Date=this.GetString("date")
	db:=Db.Create(&message)
	fmt.Println(db)
	if strings.Fields(ToString(db))[1]=="<nil>"{
		if strings.Fields(ToString(db))[2]!="0" {
			this.EchoJsonMsg("上报成功")
		}else{
			this.EchoJsonErr("上报失败")
		}

	}else
	{this.EchoJsonErr("上报失败")}

}