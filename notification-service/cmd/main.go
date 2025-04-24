package main

import (
	"log"
	"net/http"
	"notification-service/api"
	"notification-service/internal/dispatcher"
	"notification-service/internal/notifier"
	"notification-service/internal/schedular"
	"notification-service/internal/storage"
	"notification-service/internal/templates"
	"time"
)

func main() {
	notifiers := map[string]notifier.Notifier{
		"email": &notifier.EmailNotifier{},
		"sms":   &notifier.SMSNotifier{},
	}

	d := dispatcher.Dispacher{
		Notifier:       notifiers,
		TemplateEngine: &templates.SmapleTemplateEngine{},
		Storage:        storage.NewInMemoryStorage(),
		Scheduler:      &schedular.SimpleSchedular{},
	}

	data := map[string]interface{}{
		"Name":  "Sreerag",
		"Event": "Welcome",
	}
	d.Dispatch("user1", "welcome_template", data)
	d.Dispatch("user2", "welcome_template", data)

	// Wait for scheduled tasks to complete
	time.Sleep(2 * time.Second)

	// Optinal: just for upping api server
	apiHandler := &api.APIHandler{
		Dispatcher: &d,
	}

	router := api.NewRouter(apiHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
