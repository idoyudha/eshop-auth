package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idoyudha/eshop-auth/internal/usecase"
	"github.com/idoyudha/eshop-auth/pkg/logger"
)

type authUserRoutes struct {
	a usecase.AuthUser
	l logger.Interface
}

func newAuthUserRoutes(
	handler *gin.RouterGroup,
	a usecase.AuthUser,
	l logger.Interface) {

	routes := &authUserRoutes{a, l}

	h := handler.Group("/auth")
	{
		h.GET(":token", routes.getUserAuth)
	}
}

func (r *authUserRoutes) getUserAuth(ctx *gin.Context) {
	token := ctx.Param("token")
	authUser, err := r.a.GetUser(ctx, token)
	if err != nil {
		r.l.Error(err, "http - v1 - getUserAuth")
		ctx.JSON(http.StatusInternalServerError, newInternalServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, newGetSuccess(authUser))
}
