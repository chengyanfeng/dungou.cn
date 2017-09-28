package controller

import (
	"bytes"
	. "dungou.cn/datasource"
	. "dungou.cn/util"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

func (this *ApiController) Getpath() {
	sets := []Dungouset{}
	paths := make([]string, 0)
	Db.Where("status = ?", 1).Find(&sets)
	for _, v := range sets {
		path := v.Path
		paths = append(paths, path)
	}
	paths = RemoveDuplicatesAndEmpty(paths)
	this.EchoJson(paths)
}

func (this *ApiController) Getsection() {
	sets := []Dungouset{}
	path := this.GetString("path")
	sections := make([]string, 0)
	if path == "" {
		Db.Where("status = ?", 1).Find(&sets)
	} else {
		Db.Where("status = ? and path = ?", 1, path).Find(&sets)
	}
	for _, v := range sets {
		section := v.Section
		sections = append(sections, section)
	}
	sections = RemoveDuplicatesAndEmpty(sections)
	this.EchoJson(sections)
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
			updir := "upload"
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
	url := this.GetString("url")
	table := this.GetString("table")
	section := this.GetString("section")
	pg := Mysql{}

	if table == "profile" {
		profile := Profile{}
		profile.Section = section
		Db.Where("section = ?", section).Delete(Profile{})
		profile.Url = url
		Db.Create(profile)
	} else if table == "dungouset" {
		file, err := os.Open(url)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()
		reader := csv.NewReader(file)
		k := 0
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				this.EchoJsonErr(err)
				return
			}
			if k != 0 {
				inserSet(record)
			}
			k++
		}
	}else if table == "rtinfo"||table =="seclonlat"||table=="prolonlat" {
		_, e := pg.LoadCsv(url, table, ",")
		if e != nil {
			this.EchoJsonErr(e)
		}
	}
	this.EchoJson("200")

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

func inserSet(record []string) {
	set := Dungouset{}
	p := P{}

	dungou := record[3]
	status := "1"
	p["dungou"] = dungou
	p["status"] = status
	Db.Table("dungouset").Where("dungou = ? and status = ?", dungou, status).Updates(P{"status": "0"})
	set.Project = record[0]
	set.Path = record[1]
	set.Section = record[2]
	set.Dungou = record[3]
	set.Positivity = record[4]
	set.Company1 = record[5]
	set.Company2 = record[6]
	set.Client = record[7]
	set.Datano = record[8]
	set.Jack = record[9]
	set.Ringnum = record[10]
	set.Lon = record[11]
	set.Lat = record[12]
	set.Status = status
	Db.Create(set)
}
