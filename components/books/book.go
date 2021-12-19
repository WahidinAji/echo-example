package books

type Book struct {
	ID 		int 	`json:"id"`
	Author	string	`json:"author"`
	Title	string	`json:"title"`
}