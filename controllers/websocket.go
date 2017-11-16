package controllers

import (
	"chatOL/models"
	"encoding/json"
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
	this.Data["IsWebSocket"] = true
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

	//接收消息
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}
}

func broadcastWebsocket(event models.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("fail to marshal event:", err)
		return
	}

	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				unsubscribe <- sub.Value.(Subscriber).Name
			}
		}
	}
}
