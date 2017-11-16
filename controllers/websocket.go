package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	beego.Controller
}

func (this *WebSocketController) Get() {
	uname := this.GetString("uname")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}
	this.TplName = "websocket.html"
	this.Data["UserName"] = uname
}

func (this *WebSocketController) Join() {
	uname := this.GetString("uname")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, err.Error(), 400)
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	//加入聊天室
	Join(uname, ws)
	defer Leave(uname)

	//
}
