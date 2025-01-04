package entity

import "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"

type AuthUser struct {
	Username   string
	Attributes []types.AttributeType
}
