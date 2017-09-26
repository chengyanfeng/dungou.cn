package controller

import (
	//. "datahunter.cn/datasource"
	. "dungou.cn/def"
	. "dungou.cn/util"
	//"errors"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
)

type WsController struct {
	BaseController
}

func (this *WsController) Rpc() {
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		Error("Cannot setup WebSocket connection:", err)
		return
	}

	// Message receive loop.
	for {
		_, payload, err := ws.ReadMessage()
		if err != nil {
			ws.Close()
			return
		}
		fmt.Println(payload)
		//if MODE == MODE_TEST {
			Debug(payload)
		//}
		p := *JsonDecode(payload)
		body := ""
		if IsEmpty(p) {
			err = errors.New("无效的请求参数格式")
		} else {
			act := p["act"]
			args := ToP(p["args"])
			url := fmt.Sprintf("http://localhost:%v/%v", beego.BConfig.Listen.HTTPPort, act)
			header := this.Ctx.Request.Header
			hp := P{}
			for k, v := range header {
				hp[k] = v
			}
			hp["Hostname"] = this.Ctx.Request.Host
			body, err = HttpPost(url, &hp, &args)
		}
		if err != nil {
			body = JsonEncode(P{"code": GENERAL_ERR, "msg": err.Error()})
		}
		err = ws.WriteMessage(websocket.TextMessage, []byte(body))
		if err != nil {
			ws.Close()
			return
		}
	}
}
