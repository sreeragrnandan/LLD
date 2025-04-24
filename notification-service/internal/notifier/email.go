package notifier

import "fmt"

type EmailNotifier struct{}

func (e *EmailNotifier) Send(notification Notification) error {
	fmt.Printf("sending email to %s: %s", notification.UserID, notification.Message)
	return nil
}
