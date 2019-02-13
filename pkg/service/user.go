package service

import (
	"github.com/safinn/play-arch/pkg/store"
)

type UserService interface {
	Add(user *store.User) error
	Find(id int) (*store.User, error)
	FindAll() ([]*store.User, error)
	FindAllWithPets() ([]*store.User, error)
}

type userService struct {
	repo store.UserRepo
}

func NewUserService(repo store.UserRepo) UserService {
	return &userService{
		repo,
	}
}

func (s *userService) Add(user *store.User) error {
	return s.repo.Create(user)
}

func (s *userService) Find(id int) (*store.User, error) {
	return s.repo.Get(id)
}

func (s *userService) FindAll() ([]*store.User, error) {
	return s.repo.GetAll()
}

func (s *userService) FindAllWithPets() ([]*store.User, error) {
	return s.repo.GetAllWithPet()
}
