package handlers

import (
    "encoding/json"
    "net/http"
    "your-project/internal/service"
    "github.com/go-chi/chi/v5"
)

type UserHandler struct {
    userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) Routes() chi.Router {
    r := chi.NewRouter()
    r.Get("/{id}", h.GetUser)
    return r
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    // Handler implementation
}
