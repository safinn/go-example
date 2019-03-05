package store

import (
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
	Pets []*Pet `json:"pets,omitempty"` // One to Many relation
}

type userRepo struct {
	db *sqlx.DB
}

type Query struct {
	db      *sqlx.DB
	Users   []*User
	queries []func() error
}

func (q *Query) WithPets() *Query {

	queryFunc := func() error {
		var ids []int
		for _, user := range q.Users {
			ids = append(ids, user.ID)
		}

		petRepo := NewPetRepo(q.db)
		pets, error := petRepo.GetAllWithID(ids)

		if error != nil {
			return error
		}

		for _, user := range q.Users {
			for _, pet := range pets {
				if pet.UserID == user.ID {
					user.Pets = append(user.Pets, pet)
				}
			}
		}

		return nil
	}

	q.queries = append(q.queries, queryFunc)

	return q
}

func (q *Query) Exec() ([]*User, error) {
	for _, query := range q.queries {
		error := query()
		if error != nil {
			return nil, error
		}
	}

	return q.Users, nil
}

func NewUserRepository(db *sqlx.DB) *userRepo {
	return &userRepo{
		db,
	}
}

func (r *userRepo) Get(id int) *Query {
	query := &Query{
		db:      r.db,
		Users:   []*User{},
		queries: []func() error{},
	}

	queryFunc := func() error {
		user := &User{}
		error := r.db.Get(user, `SELECT * FROM "user" WHERE id = $1`, id)
		if error != nil {
			return error
		}

		query.Users = append(query.Users, user)

		return nil
	}

	query.queries = append(query.queries, queryFunc)

	return query
}

func (r *userRepo) GetAll() ([]*User, error) {
	var users []*User
	err := r.db.Select(&users, `SELECT * FROM "user"`)

	return users, err
}

func (r *userRepo) Create(user *User) error {
	stmt, _ := r.db.Prepare(`INSERT INTO "user"(id, name, age) VALUES ($1, $2, $3)`)
	stmt.Exec(user.ID, user.Name, user.Age)

	return nil
}
