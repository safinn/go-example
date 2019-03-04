package store

import (
	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	Get(id int) queryInterface
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
	WithPets() queryInterface
	Exec() []*User
}

type query struct {
	db      *sqlx.DB
	Users   []*User
	queries []func()
}

func (q *query) WithPets() queryInterface {

	queryFunc := func() {
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
	}

	q.queries = append(q.queries, queryFunc)

	return q
}

func (q *query) Exec() []*User {
	for _, query := range q.queries {
		query()
	}

	return q.Users
}

func NewUserRepository(db *sqlx.DB) UserRepo {
	return &userRepo{
		db,
	}
}

func (r *userRepo) Get(id int) queryInterface {
	query := &query{
		db:      r.db,
		Users:   []*User{},
		queries: []func(){},
	}

	queryFunc := func() {
		user := &User{}
		_ = r.db.Get(user, `SELECT * FROM user WHERE id = $1`, id)

		query.Users = append(query.Users, user)

	}

	query.queries = append(query.queries, queryFunc)

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
