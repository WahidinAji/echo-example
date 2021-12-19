package main

import (
	"echo-mysql/components/books"
	"echo-mysql/components/customer/handler"
	"echo-mysql/components/post"
	"echo-mysql/database"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)
func main()  {
	db := database.GetConn()
	defer  db.Close()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	h := handler.Handler{DB: db}
	p := post.Dependency{DB: db}
	e.GET("/customers", h.GetAll)
	e.POST("/customers", h.Create)
	e.GET("/posts", p.GetAll)
	e.GET("/posts/:id", p.GetId)
	books.BookRoute(e)
	e.Logger.Fatal(e.Start(":1323"))
}
