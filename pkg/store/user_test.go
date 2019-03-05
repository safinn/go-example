// +build integration

package store

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

func setupDb(db *sqlx.DB) {

	db.Exec(`CREATE TABLE  IF NOT EXISTS  "user" (id  integer not null primary key, name text, age integer)`)
	db.Exec(`CREATE TABLE  IF NOT EXISTS  pet (id  integer not null primary key, userid integer not null references "user"(id), name text, type text)`)

	db.Exec(`INSERT INTO "user" (id, name, age) VALUES (1, 'Dimitris', 25)`)
	db.Exec(`INSERT INTO "user" (id, name, age) VALUES (2, 'Dan', 22)`)
	db.Exec(`INSERT INTO "user" (id, name, age) VALUES (3, 'Haydn', 25)`)

	db.Exec(`INSERT INTO pet (id, userid, name, type) VALUES (1, 1, 'a', 'cat')`)
	db.Exec(`INSERT INTO pet (id, userid, name, type) VALUES (2, 3, 'b', 'cat')`)
	db.Exec(`INSERT INTO pet (id, userid, name, type) VALUES (3, 1, 'c', 'dog')`)
	db.Exec(`INSERT INTO pet (id, userid, name, type) VALUES (4, 1, 'd', 'cat')`)
}

var db *sqlx.DB

func TestMain(m *testing.M) {
	db, _ = sqlx.Open("postgres", "postgresql://riskledger:riskledger@localhost:5432/riskledger?sslmode=disable")
	setupDb(db)
	code := m.Run()
	db.Close()
	os.Exit(code)
}

func TestUserRepo_Get(t *testing.T) {
	userRepo := NewUserRepository(db)

	users, error := userRepo.Get(1).Exec()

	if error != nil {
		t.Error(error)
	}

	if len(users) != 1 {
		t.Errorf("Didn't return correct number of data")
	}
}
