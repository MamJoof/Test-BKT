package main

import "database/sql"

func openReports() (*sql.DB, error) {
	return sql.Open("postgres", "host=reports user=app dbname=reports sslmode=disable")
}
