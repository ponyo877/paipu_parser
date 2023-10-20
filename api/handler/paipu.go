package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ponyo877/paipu_parser/usecase/paipu"
)

// MakePaipuHandlers make paipu handler
func MakePaipuHandlers(e *echo.Echo, service paipu.UseCase) {
	e.GET("/paipu/parse", GetParsedPaipu(service))
}

// GetParsedPaipu get parsed paipu
func GetParsedPaipu(service paipu.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		paipu := c.QueryParam("url")
		result, err := service.GetParsedPaipu(paipu)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, result)
	}
}
