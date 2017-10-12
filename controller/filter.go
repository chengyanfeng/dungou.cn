package controller

import (
	"github.com/astaxie/beego/context"
	. "dungou.cn/def"
	. "dungou.cn/util"
	"github.com/astaxie/beego"
	"fmt"
	"errors"
	"strings"
)

var BaseFilter = func(ctx *context.Context) {
	if MODE == "test" {
		//ctx.Output.Header("Access-Control-Allow-Origin", "*")
		//ctx.Output.Header("Access-Control-Allow-Headers", "Origin,X-Requested-With,Content-Type,Accept")
		//ctx.Output.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	}
	if strings.Contains(ctx.Request.RequestURI, "?"){
		body := JsonEncode(P{"code": GENERAL_ERR, "msg": "非法访问"})
		ctx.Output.Body([]byte(body))
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
		act := p["act"].(string)
		args := ToP(p["args"])

		fmt.Println("act:"+act)
		fmt.Println("gradecookie:")
		fmt.Println(ToString(S(args["grade"].(string))))
			if 	ToString(S(args["grade"].(string)))==""&&act!="api/login" {

				err = errors.New("请先登录")
			} else {

				if ToString(S(args["grade"].(string)))!=args["grade"]&&act!="api/login"{
					err = errors.New("请先登录")

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
