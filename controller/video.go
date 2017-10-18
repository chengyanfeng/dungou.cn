package controller

import (
	"fmt"
	."dungou.cn/util"
	."dungou.cn/datasource"
	"strings"
)

type VideoController struct {
	BaseController
}
var AppKey = "5edb69eb9f5540b8aefdd41ea49002f5"
var AppSecret = "4c1ef3165da30ab01ba872c06ea9d86a"
func (this * VideoController) Update() {
	url := "https://open.ys7.com/api/lapp/token/get"
	param :=P{}
	param["appKey"]= AppKey
	param["appSecret"] = AppSecret
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
	flag := list(token)
	if flag {

	}
	this.EchoJson(200)
}

func (this * VideoController) Getsection() {
	video := []Video{}
	sections := make([]string, 0)

	Db.Find(&video)
	for _, v := range video {
		section := v.Section
		sections = append(sections, section)
	}
	sections = RemoveDuplicatesAndEmpty(sections)
	this.EchoJson(sections)
}

func (this *VideoController) Getlist() {
	section := this.GetString("section")
	video := []Video{}
	param := make(map[string]interface{})
	p :=[]P{}
	if section != "" {
		param["section"] = section
	}
	param["exception"] = "0"
	fmt.Println(param)
	Db.Where(param).Find(&video)
	fmt.Println(video)
	for _,v := range video{
		a:=P{}
		a["dungou"] = v.Dungou
		a["channelno"] = v.ChannelNo
		p = append(p,a)
	}
	this.EchoJson(p)
}

func (this *VideoController) Getvideo() {
	dungou := this.GetString("dungou")
	channelNo := this.GetString("channelNo")
	this.Require("dungou","channelNo")
	channels := strings.Split(channelNo, ",")
	videos := []Video{}
	Db.Where("dungou = ? and channel_no in (?)", dungou, channels).Find(&videos)
	this.EchoJson(videos)
}

func list(token string) bool{
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
	return getList(token,p)
}

func getList(token string,p map[string]string) bool{
	flag := false
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
		Db.Delete(&Video{})
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
			video.Dungou = deviceName
			video.Section = getsec(deviceName)
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
			flag =true
		}
	}
	return flag
}

func getsec(name string) string  {
	set := Dungouset{}
	Db.Where("dungou = ? and status = ?",name,"1").First(&set)
	section := set.Section
	return section
}