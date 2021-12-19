package handler

import (
	"database/sql"
	"echo-mysql/components/customer/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) GetAll(ctx echo.Context) error {
	db, err := h.DB.Query("SELECT id,name,email FROM customers order by id desc limit 5")
	if err != nil {
		return err
	}
	defer db.Close()
	var customers []model.Customer
	for db.Next(){
		var customer model.Customer
		err := db.Scan(&customer.ID,&customer.Name, &customer.Email)
		if err != nil {
			return err
		}
		customers = append(customers, customer)
	}
	return ctx.JSON(200, customers)
}

func (h *Handler) Create(ctx echo.Context) error {
	newCustomer := new(model.Customer)
	//newCustomer.Name = "aji echo"
	//newCustomer.Email = "aji.echo@mail.com"
	if err := ctx.Bind(newCustomer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	QUERY := "INSERT INTO customers (name, email) VALUES (?,?)"
	db, err := h.DB.ExecContext(ctx.Request().Context(),QUERY, newCustomer.Name,newCustomer.Email)
	if err != nil {
		return err
	}
	id, err := db.LastInsertId()
	if err != nil {
		return err
	}
	
	newCustomer.ID = int(id)
	return ctx.JSON(200, newCustomer)
}