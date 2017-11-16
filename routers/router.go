package routers

import (
	"chatOL/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.AppController{})
	beego.Router("/join", &controllers.AppController{}, "post:Join")
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")

}
