package post

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

//var db *sql.DB

func (d *Dependency) GetAll(ctx echo.Context) error {
	rows, err := d.FindAll(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(200,rows)
}
func (d *Dependency) GetId(ctx echo.Context) error {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}
	rows, err := d.FindId(ctx.Request().Context(), postId)
	if err != nil {
		return err
	}
	return ctx.JSON(200,rows)
}


