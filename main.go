package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	. "dungou.cn/controller"
	. "dungou.cn/datasource"
	. "dungou.cn/def"
	. "dungou.cn/task"
	. "dungou.cn/util"
	"os"

)

var orm Orm
var mssql Mssql
func main() {
	initConf()
	MODE = Trim(os.Getenv("mode"))
	beego.BConfig.Listen.HTTPPort = 9700                     //端口设置
	beego.BConfig.RecoverPanic = true                        //开启异常捕获
	beego.BConfig.EnableErrorsShow = true
	beego.BConfig.CopyRequestBody = true
	beego.InsertFilter("/*", beego.BeforeRouter, BaseFilter) //路由过滤

	//自动匹配路由
	beego.AutoRouter(&ApiController{})
	beego.InsertFilter("/rpc", beego.BeforeRouter, RpcFilter)
	//beego.InsertFilter("/api/*", beego.BeforeRouter, WhiteListFilter)
	Mkdir("./logs")                                        //创建日志文件夹
	beego.SetLogger("file", `{"filename":"logs/run.log"}`) //定义日志文件
	beego.BeeLogger.SetLogFuncCallDepth(4)
	//调用以下函数处理接口数据
	orm.Init()
	mssql.Init()

	Insert()
	go func() {
		//开启协程
		InitCache() //初始化
		crontab()   //开启定时任务
	}()
	beego.Run() //启动项目
}

func crontab() {
	toolbox.AddTask("pd", toolbox.NewTask("pd", "0 */15 * * * *", func() error {
		//每10分钟运行以下函数
		Dhq <- func() {

		}
		return nil
	}))
	toolbox.StartTask() //开启定时任务
}
func initConf() {
	myConfig := new(Config)
	config := myConfig.InitConfig("./", "privilege.ini", "nats")
	VISITOR = config["visitor"]
	VISITORS = config["visitors"]
	ORDINARYUSER = config["ordinaryuser"]
	LEADERUSER = config["leaderuser"]
	ADMINISTRATOR = config["administrator"]
	ROOT = config["root"]
	SUPERROOT = config["superroot"]



}