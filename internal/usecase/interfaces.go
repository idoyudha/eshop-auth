package usecase

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/idoyudha/eshop-auth/internal/entity"
)

type (
	AuthUser interface {
		GetUser(context.Context, string) (entity.AuthUser, error)
	}

	AuthUserWebAPI interface {
		GetUser(context.Context, *cognitoidentityprovider.GetUserInput) (*cognitoidentityprovider.GetUserOutput, error)
	}
)
