package usecase

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/idoyudha/eshop-auth/internal/entity"
)

type AuthUserUseCase struct {
	webAPI AuthUserWebAPI
}

func NewAuthUserUseCase(w AuthUserWebAPI) *AuthUserUseCase {
	return &AuthUserUseCase{
		webAPI: w,
	}
}

func (uc *AuthUserUseCase) GetUser(ctx context.Context, token string) (entity.AuthUser, error) {
	params := &cognitoidentityprovider.GetUserInput{
		AccessToken: &token,
	}

	userOutput, err := uc.webAPI.GetUser(ctx, params)
	if err != nil {
		return entity.AuthUser{}, err
	}

	authUser := entity.AuthUser{
		Username:   *userOutput.Username,
		Attributes: userOutput.UserAttributes,
	}

	return authUser, nil
}
