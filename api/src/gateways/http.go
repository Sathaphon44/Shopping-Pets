package gateways

import (
	"shoppets-api/src/services"

	"github.com/labstack/echo/v4"
)

type HttpGataways struct {
	UsersServices services.IUsersServices
	PetsServices  services.IPetsServices
	CartServices  services.ICartServices
}

func NewHTTPGateways(
	echo *echo.Echo,
	users_services services.IUsersServices,
	Pets_services services.IPetsServices,
	Cart_services services.ICartServices,
) {
	gateways := &HttpGataways{
		UsersServices: users_services,
		PetsServices:  Pets_services,
		CartServices:  Cart_services,
	}
	Routes(*gateways, echo)
}
