package service

import (
	"github.com/safinn/play-arch/pkg/store"
)

type UserRepo interface {
	Get(id int) *store.Query
	GetAll() ([]*store.User, error)
	Create(user *store.User) error
}

type UserService interface {
	Add(user *store.User) error
	Find(id int) (*store.User, error)
	FindAll() ([]*store.User, error)
	FindWithPet(id int) (*store.User, error)
}

type userService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) UserService {
	return &userService{
		repo,
	}
}

func (s *userService) Add(user *store.User) error {
	return s.repo.Create(user)
}

func (s *userService) Find(id int) (*store.User, error) {
	users, error := s.repo.Get(id).Exec()
	if error != nil {
		return nil, error
	}
	return users[0], nil
}

func (s *userService) FindAll() ([]*store.User, error) {
	return s.repo.GetAll()
}

func (s *userService) FindWithPet(id int) (*store.User, error) {
	users, error := s.repo.Get(id).WithPets().Exec()
	if error != nil {
		return nil, error
	}
	return users[0], nil
}
