package main

import (
	"database/sql"
	"net/http"
)

func sanity(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	db.Exec("DELETE FROM users WHERE id = " + r.URL.Query().Get("id"))
}
