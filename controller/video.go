package controller

import (
	"fmt"
	."dungou.cn/util"
	."dungou.cn/datasource"
)
var AppKey = "5edb69eb9f5540b8aefdd41ea49002f5"
var AppSecret = "4c1ef3165da30ab01ba872c06ea9d86a"
func GetVideo(appKey string, appSecret string) {
	url := "https://open.ys7.com/api/lapp/token/get"
	param :=P{}
	param["appKey"]= appKey
	param["appSecret"] = appSecret
	res, err := HttpPost(url,&param,&param)
	if err != nil {
		fmt.Println(err)
		return
	}
	jd := *JsonDecode([]byte(res))
	code := jd["code"].(string)
	token :=""
	if code == "200" {
		data := jd["data"].(map[string]interface{})
		token = data["accessToken"].(string)
	}
	List(token)
}

func GetList(token string,p map[string]string) {
	url :="https://open.ys7.com/api/lapp/live/video/list"
	param :=P{}
	param["accessToken"]= token
	res, err := HttpPost(url,&param,&param)
	if err != nil {
		fmt.Println(err)
	}
	jd := *JsonDecode([]byte(res))
	code := jd["code"].(string)
	if code == "200" {
		data := jd["data"].([]interface{})
		for _,v :=range data {
			video := Video{}
			a :=v.(map[string]interface{})
			deviceSerial :=a["deviceSerial"].(string)
			deviceName := p[deviceSerial]
			channelNo := a["channelNo"].(float64)
			liveAddress := a["liveAddress"].(string)
			hdAddress := a["hdAddress"].(string)
			rtmp := a["rtmp"].(string)
			rtmpHd := a["rtmpHd"].(string)
			status := a["status"].(float64)
			exception := a["exception"].(float64)
			beginTime := a["beginTime"].(float64)
			endTime := a["endTime"].(float64)

			video.DeviceSerial = deviceSerial
			video.DeviceName = deviceName
			video.ChannelNo = ToString(channelNo)
			video.LiveAddress = liveAddress
			video.HdAddress = hdAddress
			video.Rtmp = rtmp
			video.RtmpHd = rtmpHd
			video.Status = ToString(status)
			video.Exception = ToString(exception)
			video.BeginTime = ToString(beginTime)
			video.EndTime = ToString(endTime)
			Db.Create(&video)
		}
	}
}

func List(token string) {
	url :="https://open.ys7.com/api/lapp/device/list"
	p:= make(map[string]string)
	param :=P{}
	param["accessToken"]= token
	param["pageStart"] = 0
	param["pageSize"] = 50
	res, err := HttpPost(url,&param,&param)
	if err != nil {
		fmt.Println(err)
	}
	jd := *JsonDecode([]byte(res))
	code := jd["code"].(string)
	if code == "200" {
		data := jd["data"].([]interface{})
		for _,v :=range data {
			a :=v.(map[string]interface{})
			deviceName :=a["deviceName"].(string)
			deviceSerial :=a["deviceSerial"].(string)
			p[deviceSerial] = deviceName
		}
	}
	GetList(token,p)
}