package gateways

import (
	"net/http"
	"shoppets-api/domain/entities"
	"shoppets-api/src/utils"

	"github.com/labstack/echo/v4"
)

func (h HttpGataways) UpdateAndInsertCart(c echo.Context) (err error) {
	authHeader := c.Request().Header.Get("Authorization")
	token := utils.ExtractTokenFromHeader(authHeader)
	list_data, err := h.UsersServices.VerifyTokenService(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	body := new(entities.Cart)
	if err = c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	data, err := h.CartServices.UpdateCartServices(list_data.Uid.Hex(), body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

func (h HttpGataways) ShowPetsInCart(c echo.Context) (err error) {
	authHeader := c.Request().Header.Get("Authorization")
	token := utils.ExtractTokenFromHeader(authHeader)
	list_data, err := h.UsersServices.VerifyTokenService(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	statusGetDetailCart, err := h.CartServices.GetDetailCartServices(list_data.Uid.Hex())
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"data": nil,
			})
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": statusGetDetailCart,
	})
}

func (h HttpGataways) DeleteCart(c echo.Context) (err error) {
	authHeader := c.Request().Header.Get("Authorization")
	token := utils.ExtractTokenFromHeader(authHeader)
	list_data, err := h.UsersServices.VerifyTokenService(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	statusDeleteCart, err := h.CartServices.DeleteCartServices(list_data.Uid.Hex())
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": statusDeleteCart,
	})
}
