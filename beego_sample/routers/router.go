package routers

import (
	"github.com/t-fukui/beego_sample/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
