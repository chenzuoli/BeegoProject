package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 导入驱动包
	"log"
)

type Db struct {
	DbName  string
	Charset string
}

func init() {
	driverName := beego.AppConfig.String("driverName")

	// 注册驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)

	// 数据库连接配置
	mysqlUser := beego.AppConfig.String("mysqlUser")
	mysqlPwd := beego.AppConfig.String("mysqlPwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbName := beego.AppConfig.String("dbName")

	// 数据库连接url
	dbConn := mysqlUser + ":" + mysqlPwd + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8"

	err := orm.RegisterDataBase("default", driverName, dbConn)

	if err != nil {
		log.Println(err)
		log.Fatal("数据库连接失败")
	} else {
		log.Println("数据库连接成功")
	}
}

func GetDbName(db Db) string {
	return db.DbName
}

func NewDb(name string) Db {
	return Db{name, "utf8"}
}
