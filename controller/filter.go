package controller

import (
	"github.com/astaxie/beego/context"
	. "dungou.cn/def"
	. "dungou.cn/util"
	"github.com/astaxie/beego"
	"fmt"
	"errors"
)


var BaseFilter = func(ctx *context.Context) {
	if MODE == "test" {
		//ctx.Output.Header("Access-Control-Allow-Origin", "*")
		//ctx.Output.Header("Access-Control-Allow-Headers", "Origin,X-Requested-With,Content-Type,Accept")
		//ctx.Output.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	}
	fmt.Println("url:"+ctx.Request.RequestURI)
	fmt.Println("session1:")

	fmt.Println(ctx.Input.Cookie("beegosessionID"))
	//a:=ctx.Input.CruSession().SessionID()
	//fmt.Println(a)
	  ctx.Input.Cookie("beegosessionID")
	 aa:=ctx.Input.Session("beegosessionID")
	fmt.Println(aa)
	fmt.Println("ok:")
	fmt.Println(ctx.Input.CruSession.Get("username"))
	/*if ctx.Request.RequestURI!="/rpc"&&ctx.Request.RequestURI!="/api/login"{
		_, ok := ctx.Input.Session("uid").(int)

		fmt.Println(ok)
		fmt.Println("ok:")
		if !ok{
		ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		ctx.WriteString(JsonEncode(P{"code": GENERAL_ERR, "msg": "请先登录1"}))

	}
	}*/
	Debug("BaseFilter", ctx.Request.RequestURI)
}

var RpcFilter = func(ctx *context.Context) {
	fmt.Println("ctx",ctx.Input)
	p := *JsonDecode(ctx.Input.RequestBody)
	fmt.Println("p:")
	fmt.Println(p)
	fmt.Println("cooki:")
	fmt.Println(ctx.GetCookie("username"))

	var err error
	body := ""
	if IsEmpty(p) {
		err = errors.New("无效的请求参数格式")
	} else {
		fmt.Println(111111111111111)

		act := p["act"].(string)
		args := ToP(p["args"])

		fmt.Println("act；"+act)
			if IsEmpty(ctx.Input.CruSession.Get("username"))&&act!="api/login" {
				fmt.Println("session2:")
				fmt.Print(ctx.Input.CruSession.Get("username"))
				err = errors.New("请先登录2")
			} else {

				if IsEmpty(ctx.Input.CruSession.Get("username"))&&act!="api/login"{
					err = errors.New("请先登录3")

				} else {
					url := fmt.Sprintf("http://localhost:%v/%v", beego.BConfig.Listen.HTTPPort, act)
					header := ctx.Request.Header
					hp := P{}
					for k, v := range header {
						hp[k] = v
					}
					hp["Hostname"] = ctx.Request.Host
					body, err = HttpPost(url, &hp, &args)
				}
			}

	}
	if err != nil {
		body = JsonEncode(P{"code": GENERAL_ERR, "msg": err.Error()})
	} else if !IsJson([]byte(body)) {
		body = JsonEncode(P{"code": GENERAL_ERR, "msg": body})
	}
	ctx.Output.Body([]byte(body))
}
func shior(role string)(api string){

	switch role {
	case "visitor":api=VISITOR
	case "visitors":api=VISITORS
	case "ordinaryuser":api=ORDINARYUSER
	case "leaderuser":api=LEADERUSER
	case "administrator":api=ADMINISTRATOR
	case "root":api=ROOT
	case "superroot":api=SUPERROOT
	}

	return api
}