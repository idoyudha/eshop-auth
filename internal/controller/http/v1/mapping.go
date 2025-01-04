package v1

import (
	"github.com/google/uuid"
	"github.com/idoyudha/eshop-auth/internal/entity"
)

const defaultRole = "customer"

func authUserToAuthResponse(authUser entity.AuthUser) authResponse {
	response := authResponse{}
	for _, attr := range authUser.Attributes {
		switch *attr.Name {
		case "sub":
			response.UserID = uuid.MustParse(*attr.Value)
		case "custom:role":
			response.Role = *attr.Value
		}
	}
	if response.Role == "" {
		response.Role = defaultRole
	}

	return response
}
