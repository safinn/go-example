package store

import (
	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	Get(id int) *query
	GetAll() ([]*User, error)
	Create(user *User) error
}

type User struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
	Pets []*Pet `json:"pets,omitempty"` // One to Many relation
}

type userRepo struct {
	db *sqlx.DB
}

type queryInterface interface {
	WithPets() *query
}

type query struct {
	db    *sqlx.DB
	Users []*User
}

func (q *query) WithPets() *query {
	var ids []int
	for _, user := range q.Users {
		ids = append(ids, user.ID)
	}

	petRepo := NewPetRepo(q.db)
	pets, _ := petRepo.GetAllWithID(ids)

	for _, user := range q.Users {
		for _, pet := range pets {
			if pet.UserID == user.ID {
				user.Pets = append(user.Pets, pet)
			}
		}
	}

	return q
}

func NewUserRepository(db *sqlx.DB) UserRepo {
	return &userRepo{
		db,
	}
}

func (r *userRepo) Get(id int) *query {
	user := &User{}
	_ = r.db.Get(user, `SELECT * FROM user WHERE id = $1`, id)

	query := &query{
		db:    r.db,
		Users: []*User{user},
	}

	return query
}

func (r *userRepo) GetAll() ([]*User, error) {
	var users []*User
	err := r.db.Select(&users, "SELECT * FROM user")

	return users, err
}

func (r *userRepo) Create(user *User) error {
	stmt, _ := r.db.Prepare(`INSERT INTO "user"(id, name, age) VALUES ($1, $2, $3)`)
	stmt.Exec(user.ID, user.Name, user.Age)

	return nil
}
