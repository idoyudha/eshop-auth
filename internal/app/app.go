package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/idoyudha/eshop-auth/config"
	v1 "github.com/idoyudha/eshop-auth/internal/controller/http/v1"
	"github.com/idoyudha/eshop-auth/internal/usecase"
	"github.com/idoyudha/eshop-auth/internal/usecase/webapi"
	"github.com/idoyudha/eshop-auth/pkg/cognito"
	"github.com/idoyudha/eshop-auth/pkg/httpserver"
	"github.com/idoyudha/eshop-auth/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	cognitoClient := cognito.NewCognitoClient(cfg.AWSCognito)

	authUserUseCase := usecase.NewAuthUserUseCase(
		webapi.NewAuthCognitoWebAPI(cognitoClient),
	)

	handler := gin.Default()
	v1.NewRouter(handler, authUserUseCase, l)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: %s", s.String())
	case err := <-httpServer.Notify():
		l.Error("app - Run - httpServer.Notify: ", err)
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		l.Info("app - Run - httpServer.Shutdown: %s", err)
	}
}
