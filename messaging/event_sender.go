package messaging

type EventSender interface {
	SendEvent(event interface{}) error
}
