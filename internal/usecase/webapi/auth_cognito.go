package webapi

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/idoyudha/eshop-auth/pkg/cognito"
)

type AuthCognito struct {
	c cognito.CognitoClient
}

func (a *AuthCognito) GetUser(ctx context.Context, params *cognitoidentityprovider.GetUserInput) (*cognitoidentityprovider.GetUserOutput, error) {
	user, err := a.c.Client.GetUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return user, nil
}
