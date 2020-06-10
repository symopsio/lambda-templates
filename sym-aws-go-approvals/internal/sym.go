package internal

import (
	"fmt"

	"github.com/symopsio/types/go/enums"
	"github.com/symopsio/types/go/models"
)

// GetIdentity is a utility function to get a user identity for a given service
func GetIdentity(user *models.User, service enums.Service) (*models.User_Identity, error) {
	for _, identity := range user.Identities {
		if identity.Service == service {
			return identity, nil
		}
	}
	return nil, fmt.Errorf("No identity for service: %s", service)
}
