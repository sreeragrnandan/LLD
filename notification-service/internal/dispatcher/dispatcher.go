package dispatcher

import (
	"fmt"
	"notification-service/internal/notifier"
	"notification-service/internal/schedular"
	"notification-service/internal/storage"
	"notification-service/internal/templates"
)

type Dispacher struct {
	Notifier       map[string]notifier.Notifier
	TemplateEngine templates.TemplateEngine
	Storage        storage.Storage
	Scheduler      schedular.Schedular
}

func (d *Dispacher) Dispatch(userID string, templateID string, data map[string]interface{}) error {
	prefs, err := d.Storage.GetUserPreferences(userID)
	if err != nil {
		return err
	}

	message, err := d.TemplateEngine.Render(templateID, data)
	if err != nil {
		return nil
	}

	for _, channels := range prefs.Channels {
		n, ok := d.Notifier[channels]
		if !ok {
			fmt.Println("No channel found")
			continue
		}
		notification := notifier.Notification{
			UserID:  userID,
			Message: message,
		}
		err := d.Scheduler.Schedule(func() {
			err = n.Send(notification)
			if err != nil {
				fmt.Println("Error sending notification")
			}
		})
		if err != nil {
			fmt.Printf("error while scheduing notification %v", err)
		}
	}

	return nil
}
