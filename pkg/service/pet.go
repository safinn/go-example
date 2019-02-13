package service

import (
	"github.com/safinn/play-arch/pkg/store/pet"
)

type PetService interface {
	Add(pet *pet.Pet) error
	Find(id int) (*pet.Pet, error)
	FindAll() ([]*pet.Pet, error)
	FindAllWithID(ids []int) ([]*pet.Pet, error)
}

type petService struct {
	repo pet.Repo
}

func NewPetService(repo pet.Repo) PetService {
	return &petService{
		repo,
	}
}

func (s *petService) Add(pet *pet.Pet) error {
	return s.repo.Create(pet)
}

func (s *petService) Find(id int) (*pet.Pet, error) {
	return s.repo.Get(id)
}

func (s *petService) FindAll() ([]*pet.Pet, error) {
	return s.repo.GetAll()
}

func (s *petService) FindAllWithID(ids []int) ([]*pet.Pet, error) {
	return s.repo.GetAllWithID(ids)
}
