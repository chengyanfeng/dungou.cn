package controller

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
	"github.com/astaxie/beego/utils/captcha"
	. "dungou.cn/def"
	. "dungou.cn/util"
)

var Num = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Echo(msg ...interface{}) {
	var out string = ""
	for _, v := range msg {
		out += fmt.Sprintf("%v", v)
	}
	this.Ctx.WriteString(out)
}

func (this *BaseController) EchoJson(m interface{}) {
	//this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	//this.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type")
	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.WriteString(JsonEncode(P{"code": 200, "msg": m}))
}

func (this *BaseController) EchoErr(m interface{}, n string) {
	//this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	//this.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type")
	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.WriteString(JsonEncode(P{"code": 400, "msg": m, "sql":n}))
}

func (this *BaseController) EchoJsonMsg(msg interface{}) {
	//this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	//this.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type")
	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.WriteString(JsonEncode(P{"code": 200, "msg": msg}))
}

func (this *BaseController) EchoJsonOk(msg ...interface{}) {
	if msg == nil {
		msg = []interface{}{"ok"}
	}
	//this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	//this.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type")
	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.WriteString(JsonEncode(P{"code": 200, "msg": msg[0]}))
}

func (this *BaseController) EchoJsonErr(msg ...interface{}) {
	out := ""
	if msg != nil {
		for _, v := range msg {
			out = JoinStr(out, v)
		}
	}
	//this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	//this.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type")
	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.WriteString(JsonEncode(P{"code": GENERAL_ERR, "msg": out}))
}

// 把form表单内容赋予P结构体
func (this *BaseController) FormToP(keys ...string) (p P) {
	p = P{}
	r := this.Ctx.Request
	r.ParseForm()
	fmt.Println(r)
	for k, v := range r.Form {
		if len(keys) > 0 {
			if InArray(k, keys) {
				setKv(p, k, v)
			}
		} else {
			setKv(p, k, v)
		}
	}
	delete(p, "auth")
	return
}

func setKv(p P, k string, v []string) {
	if len(v) == 1 {
		if len(v[0]) > 0 {
			p[k] = v[0]
		}
	} else {
		p[k] = v
	}
}

func (this *BaseController) Captcha() {
	bytes := utils.RandomCreateBytes(4, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	str := ""
	for _, b := range bytes {
		str += Num[b]
	}
	Debug("captcha", str)
	this.SetSession("captcha", str)
	img := captcha.NewImage(bytes, 200, 80)
	if _, err := img.WriteTo(this.Ctx.ResponseWriter); err != nil {
		beego.Error("Write Captcha Image Error:", err)
	}
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func (this *BaseController) QueryParam(args ...interface{}) *P {
	p := P{}
	for _, v := range args {
		switch v.(type) {
		case string:
			k := ToString(v)
			if this.GetString(k) != "" {
				p[k] = this.GetString(k)
			}
		case P:
		// todo
		}
	}
	return &p
}

func (this *BaseController) PageParam() (start int, rows int) {
	page, _ := this.GetInt("page", 1)
	rows, _ = this.GetInt("rows", 10)
	start = (page - 1) * rows
	return
}

func (this *BaseController) Hostname() string {
	return this.Ctx.Request.Host
}

func (this *BaseController) Require(k ...string) {
	for _, v := range k {
		if IsEmpty(this.GetString(v)) {
			this.EchoJsonErr(fmt.Sprintf("需要%v参数", v))
			this.StopRun()
		}
	}
}

