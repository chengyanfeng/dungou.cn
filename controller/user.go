package controller

import (
	."dungou.cn/def"
	. "dungou.cn/datasource"
	. "dungou.cn/util"
	"fmt"
	"time"
	"strings"
)
type UserController struct {
	BaseController
}

func (this *UserController) Login() {
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

func (this *UserController)Exit(){
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
func (this *UserController)Adduser(){
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
func (this *UserController)Updateuser(){
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
func (this *UserController) Updatepassword(){
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
func (this *UserController)Finduser(){

	role := this.GetString("role")
	if role =="1"||role=="2"{
		var db interface{}
		users:=[]User{}
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

	}else {
		this.EchoJsonMsg("您无权限")
	}
}

//删除
func (this * UserController)Deletuser(){
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

//上报信息
func (this *UserController) Upmessage() {
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
func (this *UserController) Findmessage(){

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