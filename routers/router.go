package routers

import (
	. "BeegoProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &MainController{})
	beego.Router("/desc", &MysqlController{})
}
