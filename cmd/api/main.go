package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/safinn/play-arch/pkg/handlers"
	"github.com/safinn/play-arch/pkg/service"
	"github.com/safinn/play-arch/pkg/store"
	"log"
	"net/http"
)

func main() {
	// Connect to DB
	db := setupDb()

	// Abstract creation of domain layers
	userRepo := store.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	http.HandleFunc("/", userHandler.Get)
	http.HandleFunc("/id", userHandler.GetById)
	http.HandleFunc("/create", userHandler.Create)
	http.HandleFunc("/withpets", userHandler.GetWithPets)

	petRepo := store.NewPetRepo(db)
	petService := service.NewPetService(petRepo)
	petHandler := handlers.NewPetHandler(petService)

	http.HandleFunc("/pet", petHandler.Get)
	http.HandleFunc("/pet/withuser", petHandler.GetWithUser)
	http.HandleFunc("/pet/id", petHandler.GetById)
	http.HandleFunc("/pet/create", petHandler.Create)

	log.Print("Listening on 9090")
	http.ListenAndServe(":9090", nil)
}

func setupDb() *sqlx.DB {

	db, _ := sqlx.Open("sqlite3", "./foo.db")

	db.Exec(`CREATE TABLE  IF NOT EXISTS  "user" (id  integer not null primary key, name text, age integer)`)
	db.Exec(`CREATE TABLE  IF NOT EXISTS  pet (id  integer not null primary key, userid integer not null, name text, type string, FOREIGN KEY(userid) REFERENCES user(id))`)

	db.Exec(`INSERT INTO "user" (id, name, age) VALUES (1, "Dimitris", 25)`)
	db.Exec(`INSERT INTO "user" (id, name, age) VALUES (2, "Dan", 22)`)
	db.Exec(`INSERT INTO "user" (id, name, age) VALUES (3, "Haydn", 25)`)

	db.Exec(`INSERT INTO pet (id, userid, name, type) VALUES (1, 1, "a", "cat")`)
	db.Exec(`INSERT INTO pet (id, userid, name, type) VALUES (2, 3, "b", "cat")`)
	db.Exec(`INSERT INTO pet (id, userid, name, type) VALUES (3, 1, "c", "dog")`)
	db.Exec(`INSERT INTO pet (id, userid, name, type) VALUES (4, 1, "d", "cat")`)

	return db
}
