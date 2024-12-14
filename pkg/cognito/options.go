package cognito

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/idoyudha/eshop-auth/config"
)

func CognitoOptions(cognito config.AWSCognito) *cognitoidentityprovider.Client {
	options := &cognitoidentityprovider.Options{
		AppID:  cognito.AppID,
		Region: cognito.Region,
	}

	return cognitoidentityprovider.New(*options)
}
