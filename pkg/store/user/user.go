package user

import (
	"github.com/jmoiron/sqlx"
)

type Repo interface {
	Get(id int) (*User, error)
	GetAll() ([]*User, error)
	Create(user *User) error
}

type User struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) Repo {
	return &userRepo{
		db,
	}
}

func (r *userRepo) Get(id int) (*User, error) {
	user := &User{}
	err := r.db.Get(user, `SELECT * FROM user WHERE id = $1`, id)

	return user, err
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
