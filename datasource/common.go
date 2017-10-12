package datasource

import (
	. "dungou.cn/util"
	"strings"
)

var Conn = P{
	/*JY正式*/
	"username": "root",
	"password": "root",
	"host":     "localhost",
	"port":     3306,
	"name":     "test",
	/*JY备份*/
	//"username": "sltx",
	//"password": "sltx@2017",
	//"host":     "172.17.129.11",
	//"port":     3306,
	//"name":     "sltx",
}

var SqlServerConn = P{
	/*JY正式*/
	"username": "sa",
	"password": "sa63189188sa",
	"host":     "116.228.163.189",
	"name":     "dtdg",
	"port":     1433,
	/*JY备份*/
	//"username": "sa",
	//"password": "datahunter",
	//"host":     "localhost",
	//"port":     1433,
	//"name":     "dtdg",
}

var TjConn = P{
	"username": "sa",
	"password": "datahunter",
	"host":     "localhost",
	"port":     1433,
	"name":     "SMCAS",
}

func FilterAs(o string) (n string) {
	tmp := strings.Split(o, " as ")
	if len(tmp) > 1 {
		n = tmp[0]
	} else {
		n = o
	}
	return
}
