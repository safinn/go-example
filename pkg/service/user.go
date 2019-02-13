package service

import "github.com/safinn/play-arch/pkg/store/user"

type UserService interface {
	Add(user *user.User) error
	Find(id int) (*user.User, error)
	FindAll() ([]*user.User, error)
}

type userService struct {
	repo user.Repo
}

func NewUserService(repo user.Repo) UserService {
	return &userService{
		repo,
	}
}

func (s *userService) Add(user *user.User) error {
	return s.repo.Create(user)
}

func (s *userService) Find(id int) (*user.User, error) {
	return s.repo.Get(id)
}

func (s *userService) FindAll() ([]*user.User, error) {
	return s.repo.GetAll()
}
