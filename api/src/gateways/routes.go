package gateways

import (
	"github.com/labstack/echo/v4"
)

func Routes(gateways HttpGataways, route *echo.Echo) {
	// customer
	route.GET("/users/showusersall", gateways.Show_users_all)
	route.POST("/users/register", gateways.Register)
	route.POST("/users/login", gateways.Login)
	route.GET("/users/session", gateways.VerifyToken)
	route.DELETE("/users/deleteaccount", gateways.DeleteAccount)

	// pets
	route.GET("/pets/showpetsall", gateways.GetPetsAllProcess)

	// carts
	route.POST("/carts/update", gateways.UpdateAndInsertCart)
	route.GET("/carts/find", gateways.ShowPetsInCart)
	route.DELETE("/carts/delete", gateways.DeleteCart)
}
