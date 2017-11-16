package controllers

import "github.com/astaxie/beego"

type AppController struct {
	beego.Controller
}

//注释
func (this *AppController) Get() {
	this.TplName = "welcome.html"
}
