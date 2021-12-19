package post

import "database/sql"

type Post struct{
	ID int		`json:"id"`
	Title string	`json:"title"`
	Desc string		`json:"desc"`
}

type Dependency struct {
	DB *sql.DB
}