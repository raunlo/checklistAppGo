//go:generate go run github.com/google/wire/cmd/wire@latest
//go:build wireinject
// +build wireinject

package infra

import (
	"com.rlohmus.checklist/internal/controllers"
	controllerMapper "com.rlohmus.checklist/internal/controllers/mapper"
	"com.rlohmus.checklist/internal/repository"
	dboMapper "com.rlohmus.checklist/internal/repository/mapper"
	"com.rlohmus.checklist/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(configuration ApplicationConfiguration) Application {
	wire.Build(
		CreateApplication,
		GetGinRouter,
		CreateHttpV1Router,
		// checklist resource set
		wire.NewSet(controllerMapper.NewChecklistDtoMapper,
			controllers.CreateChecklistController,
			service.CreateChecklistService,
			repository.CreateChecklistRepository,
			dboMapper.NewChecklistDboMapper),
		// checklist item resource set
		wire.NewSet(controllers.CreateChecklistItemsController,
			service.CreateChecklistItemService,
			repository.CreateChecklistItemRepository),
		// checklist item template resource set
		wire.NewSet(controllerMapper.NewChecklistItemTemplateDtoMapper,
			controllers.CreateChecklistItemTemplateController,
			service.CreateChecklistItemTemplateService,
			repository.CreateChecklistItemTemplateRepository),
		getGORMDb,
		wire.FieldsOf(new(ApplicationConfiguration), "DatabaseConfiguration"),
		wire.FieldsOf(new(ApplicationConfiguration), "ServerConfiguration"),
	)
	return Application{}
}

func getGORMDb(configuration DatabaseConfiguration) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
		configuration.DatabaseHost, configuration.DatabaseUser, configuration.DatabasePassword, configuration.DatabaseName,
		configuration.DatabasePort, configuration.SslMode.Get())
	if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		return db
	}
}

func GetGinRouter() *gin.Engine {
	return gin.New()
}
