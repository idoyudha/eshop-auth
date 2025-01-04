package webapi

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/idoyudha/eshop-auth/pkg/cognito"
)

type AuthCognitoWebAPI struct {
	*cognito.CognitoClient
}

func NewAuthCognitoWebAPI(client *cognito.CognitoClient) *AuthCognitoWebAPI {
	return &AuthCognitoWebAPI{
		client,
	}
}

func (a *AuthCognitoWebAPI) GetUser(ctx context.Context, params *cognitoidentityprovider.GetUserInput) (*cognitoidentityprovider.GetUserOutput, error) {
	user, err := a.Client.GetUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return user, nil
}
