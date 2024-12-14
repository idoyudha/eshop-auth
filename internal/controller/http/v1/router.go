package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idoyudha/eshop-auth/internal/usecase"
	"github.com/idoyudha/eshop-auth/pkg/logger"
)

func NewRouter(
	handler *gin.Engine,
	uca usecase.AuthUser,
	l logger.Interface,
) {
	handler.Use(gin.Recovery())
	handler.Use(gin.Logger())

	handler.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	h := handler.Group("/v1")
	{
		newAuthUserRoutes(h, uca, l)
	}
}
