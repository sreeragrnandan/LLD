package notifier

type NotifType string

type Notification struct {
	UserID  string
	Message string
}

type Notifier interface {
	Send(notification Notification) error
}
