package main

import (
	"database/sql"
	"net/http"
)

func newIssue(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	db.Exec("UPDATE users SET name = '" + r.URL.Query().Get("name") + "' WHERE id = 1")
}
