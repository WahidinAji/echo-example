package books

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RepoFindAll(c echo.Context) error {
	books, err := FindAll(c.Request().Context())
	if err != nil {
		return err 
	}
	return c.JSON(http.StatusOK, books)
}