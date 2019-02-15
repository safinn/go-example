package handlers

import (
	"encoding/json"
	"github.com/safinn/play-arch/pkg/service"
	"github.com/safinn/play-arch/pkg/store"
	"net/http"
	"strconv"
)

type UserHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetWithPets(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService,
	}
}

func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	users, _ := h.userService.FindAll()

	response, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *userHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	idInt, _ := strconv.Atoi(id)
	user, _ := h.userService.Find(idInt)

	response, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u store.User
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&u)

	_ = h.userService.Add(&u)

	response, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *userHandler) GetWithPets(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	idInt, _ := strconv.Atoi(id)
	user, _ := h.userService.FindWithPet(idInt)

	response, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
