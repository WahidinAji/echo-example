package books

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

var db *sql.DB

func RepoFindAll(c echo.Context) error {
	books, err := FindAll(c.Request().Context())
	if err != nil {
		return err 
	}
	return c.JSON(http.StatusOK, books)
}