package gateways

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h HttpGataways) GetPetsAllProcess(c echo.Context) error {
	data, err := h.PetsServices.GetPetsAllServices()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}
