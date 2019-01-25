package main

import (
	_ "blockchain_explorer/routers"
	"blockchain_explorer/schedule"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
)

func init() {
	path := beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@tcp(" + beego.AppConfig.String("mysqlhost") + ":" + beego.AppConfig.String("mysqlport") + ")/" + beego.AppConfig.String("mysqldb") + "?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", path, maxIdle, maxConn)
	//orm.SetMaxIdleConns("default", 10)
	//orm.SetMaxOpenConns("default", 10)
}

func main() {
	i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")
	i18n.SetMessage("en-US", "conf/locale_en-US.ini")
	beego.AddFuncMap("i18n", i18n.Tr)
	schedule.OpenTask()
	beego.Run()
}
