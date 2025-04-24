package api

import (
	"encoding/json"
	"net/http"
	"notification-service/internal/dispatcher"

	"github.com/gorilla/mux"
)

type NotificationRequest struct {
	UserID     string                 `json:"user_id"`
	TemplateID string                 `json:"template_id"`
	Data       map[string]interface{} `json:"data"`
}

type APIHandler struct {
	Dispatcher *dispatcher.Dispacher
}

func (h *APIHandler) SendNotification(w http.ResponseWriter, r *http.Request) {
	var req NotificationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.Dispatcher.Dispatch(req.UserID, req.TemplateID, req.Data)
	if err != nil {
		http.Error(w, "Failed to dispatch notification", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification dispatched successfully"))
}

func NewRouter(handler *APIHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/send-notification", handler.SendNotification).Methods("POST")
	return router
}
