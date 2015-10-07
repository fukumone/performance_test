package controllers

import (
	"github.com/astaxie/beego"
)

type Result struct {
    Response string
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
    result := Result{Response: "OK"}
    c.Data["json"] = &result
    c.ServeJson()
}
