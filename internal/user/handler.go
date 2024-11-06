package user

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// A handler is a function or method responsible for directly processing HTTP requests and generating responses
type Handler struct {
	UserService Service
}

func NewHandler(userService Service) *Handler {
	return &Handler{
		UserService: userService,
	}
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid User Id", http.StatusBadRequest)
	}

	user, err := h.UserService.GetUserById(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
	} else {
		w.Header().Set("Content-type", "application-json")
		json.NewEncoder(w).Encode(user)
	}
}
