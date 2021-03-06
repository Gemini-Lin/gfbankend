package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var Log *logs.BeeLogger = logs.NewLogger(10000) //定义日志Log，因为需要在整个project用到，所以需要定义为全局

func init() {
	//日志暂且设置输出到控制台
	if err := Log.SetLogger("console"); err != nil {
		fmt.Println("fail to set logger!")
		os.Exit(-1)
	}
	Log.SetLevel(logs.LevelDebug)
	Log.EnableFuncCallDepth(true)
	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		Log.Error("fail to register driver", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		beego.AppConfig.String("mysql::user"),
		beego.AppConfig.String("mysql::password"),
		beego.AppConfig.String("mysql::addr"),
		beego.AppConfig.String("mysql::database"))

	if err := orm.RegisterDataBase("default", "mysql", dsn); err != nil {
		Log.Error("fail to register database", err)
	}
	orm.RegisterModel(new(Card)) //登记orm
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Enterprise))
	orm.RegisterModel(new(DelCard))
}
