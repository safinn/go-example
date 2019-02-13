package handlers

import (
	"encoding/json"
	"github.com/safinn/play-arch/pkg/service"
	"github.com/safinn/play-arch/pkg/store/pet"
	"net/http"
	"strconv"
)

type PetHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type petHandler struct {
	petService service.PetService
}

func NewPetHandler(petService service.PetService) PetHandler {
	return &petHandler{
		petService,
	}
}

func (h *petHandler) Get(w http.ResponseWriter, r *http.Request) {
	pets, _ := h.petService.FindAll()

	response, _ := json.Marshal(pets)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *petHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	idInt, _ := strconv.Atoi(id)
	pet, _ := h.petService.Find(idInt)

	response, _ := json.Marshal(pet)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *petHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p pet.Pet
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&p)

	_ = h.petService.Add(&p)

	response, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
