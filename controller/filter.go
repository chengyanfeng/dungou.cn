package controller

import (
	"github.com/astaxie/beego/context"
	. "dungou.cn/def"
	. "dungou.cn/util"
	"github.com/astaxie/beego"
	"fmt"
	"errors"
	."github.com/casbin/casbin"
)


var BaseFilter = func(ctx *context.Context) {
	if MODE == "test" {
		//ctx.Output.Header("Access-Control-Allow-Origin", "*")
		//ctx.Output.Header("Access-Control-Allow-Headers", "Origin,X-Requested-With,Content-Type,Accept")
		//ctx.Output.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	}
	Debug("BaseFilter", ctx.Request.RequestURI)
}

var RpcFilter = func(ctx *context.Context) {
	fmt.Println("ctx",ctx.Input)
	p := *JsonDecode(ctx.Input.RequestBody)
	var err error
	body := ""
	if IsEmpty(p) {
		err = errors.New("无效的请求参数格式")
	} else {
		act := p["act"]
		args := ToP(p["args"])
		role("","")
		url := fmt.Sprintf("http://localhost:%v/%v", beego.BConfig.Listen.HTTPPort, act)
		header := ctx.Request.Header
		hp := P{}
		for k, v := range header {
			hp[k] = v
		}
		hp["Hostname"] = ctx.Request.Host
		body, err = HttpPost(url, &hp, &args)
	}
	if err != nil {
		body = JsonEncode(P{"code": GENERAL_ERR, "msg": err.Error()})
	} else if !IsJson([]byte(body)) {
		body = JsonEncode(P{"code": GENERAL_ERR, "msg": body})
	}
	ctx.Output.Body([]byte(body))
}

func role(role string,api string) {
	e := NewEnforcer("path/to/rbac_model.conf", "path/to/rbac_policy.csv")
	testEnforceWithoutUsers( e, "alice", "read", true)
}

func testEnforceWithoutUsers(e *Enforcer, obj string, act string, res bool) {
	if e.Enforce(obj, act) != res {
		fmt.Println("%s, %s: %t, supposed to be %t", obj, act, !res, res)
	}
}