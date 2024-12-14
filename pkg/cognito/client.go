package cognito

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/idoyudha/eshop-auth/config"
)

type CognitoClient struct {
	Client *cognitoidentityprovider.Client
}

func NewCognitoClient(cognito config.AWSCognito) *CognitoClient {
	return &CognitoClient{
		Client: CognitoOptions(cognito),
	}
}
