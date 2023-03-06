package main

import (
	"net/http"

	"backend/src/handler"
	"backend/src/model"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Validator = &model.CustomValidator{Validator: validator.New()}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodPost},
	}))

	e.POST("/fizzbuzz", handler.GetFizzBuzz)
	e.Logger.Fatal(e.Start(":3000"))
}
