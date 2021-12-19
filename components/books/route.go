package books

import "github.com/labstack/echo/v4"

func BookRoute(app *echo.Echo)  {
	app.GET("/books", RepoFindAll)
}
