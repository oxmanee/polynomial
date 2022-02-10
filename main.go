package main

import (
	"polynomial/delivery/http"
	"polynomial/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	polynomialUsecase := usecase.NewPolynomialUsecase()

	e := echo.New()
	http.NewPolynomialDelivery(e, polynomialUsecase)

	e.Logger.Fatal(e.Start(":80"))
}
