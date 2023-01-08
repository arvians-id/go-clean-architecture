//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/arvians-id/go-clean-architecture/cmd/config"
	"github.com/arvians-id/go-clean-architecture/internal/http/controller"
	"github.com/arvians-id/go-clean-architecture/internal/http/routes"
	"github.com/arvians-id/go-clean-architecture/internal/repository"
	"github.com/arvians-id/go-clean-architecture/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
	controller.NewUserController,
)

func InitServerAPI(configuration config.Config) (*gin.Engine, error) {
	wire.Build(
		config.NewInitializedDatabase,
		userSet,
		routes.NewInitializedRoutes,
	)

	return nil, nil
}
