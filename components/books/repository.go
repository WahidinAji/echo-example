package books

import (
	"context"
	"echo-mysql/database"
)


func FindAll(ctx context.Context) ([]Book, error) {
	db := database.GetConn()
	defer db.Close()
	query := "SELECT id, author, title FROM books"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []Book
	for rows.Next(){
		var book Book
		err := rows.Scan(&book.ID, &book.Author, &book.Title)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
