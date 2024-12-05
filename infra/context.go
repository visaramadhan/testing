package infra

import (
	"go.uber.org/zap"
	"project/config"
	"project/database"
	"project/handler"
	"project/log"
	"project/repository"
	"project/service"
)

type ServiceContext struct {
	Cfg config.Config
	Ctl handler.Handler
	Log *zap.Logger
}

func NewServiceContext() (*ServiceContext, error) {

	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	config, err := config.LoadConfig()
	if err != nil {
		handlerError(err)
	}

	// instance looger
	log, err := log.InitZapLogger(config)
	if err != nil {
		handlerError(err)
	}

	// instance database
	db, err := database.ConnectDB(config)
	if err != nil {
		handlerError(err)
	}

	// instance repository
	repository := repository.NewRepository(db)

	// instance service
	service := service.NewService(repository)

	// instance controller
	Ctl := handler.NewHandler(service, log)

	return &ServiceContext{Cfg: config, Ctl: *Ctl, Log: log}, nil
}
