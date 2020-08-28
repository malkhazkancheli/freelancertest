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

	//TODO: Search for ueID in DB (Malkhaz)
	payload := &models.AuthenticationSubscription{
		AlgorithmID:                   "test AlgorithmID for:" + ueID,
		AuthenticationManagementField: "test AuthenticationManagementField for:"+ueID,
		AuthenticationMethod:          swag.String("test AuthenticationMethod for:"+ueID),
		EncOpcKey:                     "test EncOpcKey for:"+ueID,
		EncPermanentKey:               "test EncPermanentKey for:"+ueID,
		EncTopcKey:                    "test EncTopcKey for:"+ueID,
		ProtectionParameterID:         "test ProtectionParameterID for:"+ueID,
	}

	return payload,nil
}
