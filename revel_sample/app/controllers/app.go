package controllers

import "github.com/revel/revel"

type Result struct {
    Response string
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
    result := Result{Response: "OK"}
	return c.RenderJson(result)
}
