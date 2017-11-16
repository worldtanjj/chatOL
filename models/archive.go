package models

type EventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
)

type Event struct {
	Type      EventType //事件类型
	User      string    //用户
	Timestamp int       //timestamp (secs)
	Content   string    //内容
}
