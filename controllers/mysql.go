package controllers

import (
	"BeegoProject/models"
	_ "BeegoProject/models"
	"github.com/astaxie/beego"
)

type MysqlController struct {
	beego.Controller
}

func (c *MysqlController) GetDb() models.Db {
	db := models.Db{
		DbName:  "default",
		Charset: "utf8"}
	return db
}

func (c *MysqlController) Get() {
	c.Data["dbName"] = "default"
}
