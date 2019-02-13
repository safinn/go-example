package service

import (
	"github.com/safinn/play-arch/pkg/store"
)

type PetService interface {
	Add(pet *store.Pet) error
	Find(id int) (*store.Pet, error)
	FindWithUser(id int) (*store.Pet, error)
	FindAll() ([]*store.Pet, error)
	FindAllWithID(ids []int) ([]*store.Pet, error)
}

type petService struct {
	repo store.PetRepo
}

func NewPetService(repo store.PetRepo) PetService {
	return &petService{
		repo,
	}
}

func (s *petService) Add(pet *store.Pet) error {
	return s.repo.Create(pet)
}

func (s *petService) Find(id int) (*store.Pet, error) {
	return s.repo.Get(id)
}

func (s *petService) FindAll() ([]*store.Pet, error) {
	return s.repo.GetAll()
}

func (s *petService) FindAllWithID(ids []int) ([]*store.Pet, error) {
	return s.repo.GetAllWithID(ids)
}

func (s *petService) FindWithUser(id int) (*store.Pet, error) {
	return s.repo.GetWithUser(id)
}
