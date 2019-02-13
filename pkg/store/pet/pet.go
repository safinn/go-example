package pet

import (
	"github.com/jmoiron/sqlx"
	"strconv"
)

type Repo interface {
	Get(int) (*Pet, error)
	GetAll() ([]*Pet, error)
	GetAllWithID([]int) ([]*Pet, error)
	Create(pet *Pet) error
}

type Pet struct {
	ID     int    `json:"id" db:"id"`
	UserID int    `json:"user_id" db:"userid"`
	Name   string `json:"name" db:"name"`
	Type   string `json:"type" db:"type"`
}

type petRepo struct {
	db *sqlx.DB
}

func NewPetRepo(db *sqlx.DB) Repo {
	return &petRepo{
		db,
	}
}

func (r *petRepo) Get(id int) (*Pet, error) {
	pet := &Pet{}
	err := r.db.Get(pet, `SELECT * FROM pet WHERE id = $1`, id)

	return pet, err
}

func (r *petRepo) GetAll() ([]*Pet, error) {
	var pets []*Pet
	err := r.db.Select(&pets, "SELECT * FROM pet")

	return pets, err
}

func (r *petRepo) Create(pet *Pet) error {
	stmt, _ := r.db.Prepare(`INSERT INTO pet(id, userid, name, type) VALUES ($1, $2, $3, $4)`)
	stmt.Exec(pet.ID, pet.UserID, pet.Name, pet.Type)

	return nil
}

func (r *petRepo) GetAllWithID(ids []int) ([]*Pet, error) {
	var pets []*Pet

	conditions := ""
	for i, id := range ids {
		if i != 0 {
			conditions += " OR "
		}

		conditions += "userid = " + strconv.Itoa(id)
	}

	err := r.db.Select(&pets, "SELECT * FROM pet WHERE "+conditions)

	return pets, err
}
