package routers

import (
	"Beego9.24/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/register", &controllers.MainController{})
}
