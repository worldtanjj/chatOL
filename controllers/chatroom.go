package controllers

import (
	"chatOL/models"

	"github.com/gorilla/websocket"
)

type Subscription struct {
	Archive []models.Event      //所有档案
	New     <-chan models.Event //新事件
}
type Subscriber struct {
	Name string
	Conn *websocket.Conn
}

var (
	//Channel for new join users
	subscribe = make(chan Subscriber, 10)

	//Channel for exit users
	unsubscribe = make(chan string, 10)

	//Send Event here to publish them
	publish = make(chan models.Event, 10)
)

func Join(uname string, ws *websocket.Conn) {
	subscribe <- Subscriber{Name: uname, Conn: ws}
}
func Leave(user string) {
	unsubscribe <- user
}
