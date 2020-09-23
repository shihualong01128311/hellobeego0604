package routers

import (
	"BeegoHello0604/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/index", &controllers.MainController{})
}
