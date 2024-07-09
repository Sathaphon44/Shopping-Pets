package main

import (
	"log"
	"shoppets-api/domain/datasources"
	"shoppets-api/domain/repository"
	"shoppets-api/src/gateways"
	"shoppets-api/src/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/radulucut/dotenv"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	err := dotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set database mongodb
	db := datasources.ConnectMongoDB()

	// Set new Customer Repository.
	repo := repository.NewUserRepo(db)
	customerServices := services.NewUsersServices(repo)

	// Set new Pets Repository.
	PetsRepo := repository.NewPetsRepo(db)
	petsServices := services.NewPetsServiecs(PetsRepo)

	// Set new Carts Repository.
	CartsRepo := repository.NewCartRepo(db)
	CartsServices := services.NewCartServices(CartsRepo)

	// Set new Gateways.
	gateways.NewHTTPGateways(
		e,
		customerServices,
		petsServices,
		CartsServices,
	)

	e.Logger.Fatal(e.Start(":8080"))
}
