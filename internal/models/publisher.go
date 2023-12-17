package models

const PUBLISHER_SCHEMA = `
CREATE TABLE IF NOT EXISTS publishers (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	first_name text NOT NULL,
	last_name text NOT NULL);
`

type Publisher struct {
	Id        int    `json:"id" uri:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}
