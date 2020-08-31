package auth

import (
	"freelancertest/models"

	"github.com/go-openapi/swag"
)

// AuthenticationSubscriptionService is an implementation of service working with the data
type AuthenticationSubscriptionService struct {
}

// NewAuthenticationSubscriptionService constructor
func NewAuthenticationSubscriptionService() *AuthenticationSubscriptionService {

	return &AuthenticationSubscriptionService{}
}

// AuthenticationSubscriptionSearch returns object by it's UeID
func (s *AuthenticationSubscriptionService) AuthenticationSubscriptionSearch(ueID string) (*models.AuthenticationSubscription, *models.ProblemDetails) {

	// AuthenticationSubscription data structure based on swagger file
	// properties:
	//   authenticationMethod:
	//     type: string
	//     enum:
	//       - 5G_AKA
	//       - EAP_AKA_PRIME
	//   encPermanentKey:
	//     type: string
	//   protectionParameterId:
	//     type: string
	//   authenticationManagementField:
	//     type: string
	//     pattern: ^[A-Fa-f0-9]{4}$
	//   algorithmId:
	//     type: string
	//   encOpcKey:
	//     type: string
	//   encTopcKey:
	//     type: string
	// required:
	// - authenticationMethod

	//TODO: Search for ueID in DB (Malkhaz)
	payload := &models.AuthenticationSubscription{
		AlgorithmID:                   "123456",
		AuthenticationManagementField: "A4E3",                 //pattern: ^[A-Fa-f0-9]{4}$
		AuthenticationMethod:          swag.String("5G_AKA:"), //enum: 5G_AKA/EAP_AKA_PRIME
		EncOpcKey:                     "encoptkey",
		EncPermanentKey:               "encpermanentkey",
		EncTopcKey:                    "enctopckey",
		ProtectionParameterID:         "protectionparameterid",
	}

	return payload, nil
}
