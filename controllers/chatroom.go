package controllers

import (
	"chatOL/models"
	"container/list"
	"time"

	"github.com/astaxie/beego"

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
	publish     = make(chan models.Event, 10)
	subscribers = list.New()
)

func init() {
	go chatroom()
}

func Join(uname string, ws *websocket.Conn) {
	subscribe <- Subscriber{Name: uname, Conn: ws}
}

func Leave(user string) {
	unsubscribe <- user
}
func chatroom() {
	for {
		select {
		case sub := <-subscribe:
			if !isUserExist(subscribers, sub.Name) {
				subscribers.PushBack(sub)
				publish <- newEvent(models.EVENT_JOIN, sub.Name, "")
				beego.Info("New User:", sub.Name)
			} else {
				beego.Info("Old User:", sub.Name)
			}
		case unsub := <-unsubscribe:
			for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Name == unsub {
					subscribers.Remove(sub)
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
						beego.Error("websocket closed :", unsub)
					}
					publish <- newEvent(models.EVENT_LEAVE, unsub, "")
					break
				}
			}
		case event := <-publish:
			broadcastWebsocket(event)
			if event.Type == models.EVENT_MESSAGE {
				beego.Info("Message from ", event.User, ";Content:", event.Content)
			}
		}
	}
}

func newEvent(ep models.EventType, user, msg string) models.Event {
	return models.Event{Type: ep, User: user, Timestamp: int(time.Now().Unix()), Content: msg}
}

func isUserExist(subscribers *list.List, user string) bool {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).Name == user {
			return true
		}
	}
	return false
}
