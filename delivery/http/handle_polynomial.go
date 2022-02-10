package http

import (
	"net/http"
	"polynomial/constant"
	"polynomial/model"
	usecase "polynomial/usecase/interface"

	validate "polynomial/validator"

	"github.com/labstack/echo/v4"
)

type polynomialDelivery struct {
	polynomialUsecase usecase.PolynomialUsecaseInterface
}

func NewPolynomialDelivery(e *echo.Echo, polynomialUsecase usecase.PolynomialUsecaseInterface) {
	handler := &polynomialDelivery{
		polynomialUsecase: polynomialUsecase,
	}

	polynomialRoute := e.Group("/polynomial")

	polynomialRoute.GET("/dataset", handler.Dataset)
	polynomialRoute.POST("/calculate", handler.Calculate)
}

func (handler *polynomialDelivery) Calculate(c echo.Context) error {

	var response model.CalculateResponse

	request := new(model.CalculateRequest)

	err := c.Bind(request)
	if err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	err = validate.CheckRequest(*request)
	if err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	calculateRes := handler.polynomialUsecase.Calculate(handler.polynomialUsecase.PrepareDataset(*request), 0)

	checkDataset := handler.polynomialUsecase.CheckDataset(calculateRes)
	if checkDataset {
		response.Message = constant.SUCCESS
	} else {
		response.Message = constant.FAIL
	}

	return c.JSON(http.StatusOK, response)
}

func (handler *polynomialDelivery) Dataset(c echo.Context) error {

	var response model.CalculateResponse

	response.Dataset = handler.polynomialUsecase.RadomArg()

	return c.JSON(http.StatusOK, response)
}
