package main

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "host=localhost user=app dbname=app sslmode=disable")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/users", getUsers)
	http.ListenAndServe(":8080", nil)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT name FROM users")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()
}
// docs
