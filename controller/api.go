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
