package controllers

import "github.com/astaxie/beego"

type AppController struct {
	beego.Controller
}

//Get 进入界面
func (this *AppController) Get() {
	// this.Ctx.Output.Body([]byte("你好"))
	this.TplName = "welcome.html"
}

//Join 处理 post请求
func (this *AppController) Join() {
	uname := this.GetString("uname")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	this.Redirect("ws?uname="+uname, 302)
	return
}
