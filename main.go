package main

import (
	"pybbs-go/models"
	_ "pybbs-go/routers"
	_ "pybbs-go/templates"
	_ "pybbs-go/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	url := beego.AppConfig.String("jdbc.url")
	port := beego.AppConfig.String("jdbc.port")
	username := beego.AppConfig.String("jdbc.username")
	password := beego.AppConfig.String("jdbc.password")
	orm.RegisterDataBase("default", "mysql", username+":"+password+"@tcp("+url+":"+port+")/pybbs-go?charset=utf8mb4&parseTime=true&charset=utf8mb4&loc=Asia%2FShanghai", 30)
	orm.RegisterModel(
		new(models.User),
		new(models.Topic),
		new(models.Section),
		new(models.Reply),
		new(models.ReplyUpLog),
		new(models.Role),
		new(models.Permission))
	orm.RunSyncdb("default", false, true)
}

func main() {
	//orm.Debug = true
	//ok, err := regexp.MatchString("/topic/edit/[0-9]+", "/topic/edit/123")
	//beego.Debug(ok, err)
	beego.SetLogger("file", `{"filename":"logs/staging.log"}`)
	// beego.BeeLogger.DelLogger("console")

	beego.Run()
}
