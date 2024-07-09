package gateways

import (
	"net/http"
	"shoppets-api/domain/entities"
	"shoppets-api/src/utils"

	"github.com/labstack/echo/v4"
)

func (h HttpGataways) Show_users_all(c echo.Context) error {
	data, err := h.UsersServices.Show_users_all()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

func (h HttpGataways) Register(c echo.Context) (err error) {
	data := new(entities.RegistrationUsers)
	if err = c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	insert_status, err := h.UsersServices.RegisterService(data.Username, data.Email, data.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if insert_status == int64(1) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "ไม่สามารถเพิ่มบัญชีได้เนื่องจากมีบัญชี email นี้อยู่ในระบบแล้ว",
		})
	}
	return c.JSON(http.StatusOK, insert_status)
}

func (h HttpGataways) Login(c echo.Context) (err error) {
	data := new(entities.RegistrationUsers)
	if err = c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	login_status, err := h.UsersServices.LoginServices(data.Email, data.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": login_status,
		})
	}
	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"token": login_status,
	})
}

func (h HttpGataways) VerifyToken(c echo.Context) (err error) {
	authHeader := c.Request().Header.Get("Authorization")
	token := utils.ExtractTokenFromHeader(authHeader)
	result, err := h.UsersServices.VerifyTokenService(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"userId": result.Uid, "email": result.Email, "username": result.Username})
}

func (h HttpGataways) DeleteAccount(c echo.Context) (err error) {
	authHeader := c.Request().Header.Get("Authorization")
	token := utils.ExtractTokenFromHeader(authHeader)
	status_delete, err := h.UsersServices.DeleteAccountService(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, status_delete)
}
