package main

import (
	"github.com/arvians-id/go-clean-architecture/cmd/config"
	"github.com/arvians-id/go-clean-architecture/internal/http/controller"
	"github.com/arvians-id/go-clean-architecture/internal/repository"
	"github.com/arvians-id/go-clean-architecture/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	configuration := config.New("../../.env")

	router := gin.Default()
	db, err := config.NewInitializedDatabase(configuration)
	if err != nil {
		log.Fatalln(err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userController.Route(router)

	err = router.Run(configuration.Get("APP_PORT"))
	if err != nil {
		log.Fatalln(err)
	}
}
