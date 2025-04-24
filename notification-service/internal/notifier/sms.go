package notifier

import "fmt"

type SMSNotifier struct{}

func (e *SMSNotifier) Send(notification Notification) error {
	fmt.Printf("sending SMS to %s: %s", notification.UserID, notification.Message)
	return nil
}
