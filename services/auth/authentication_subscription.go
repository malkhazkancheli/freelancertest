package auth

import "freelancertest/models"

// AuthenticationSubscriptionService is an implementation of service working with the data
type AuthenticationSubscriptionService struct {

}

// NewAuthenticationSubscriptionService constructor
func NewAuthenticationSubscriptionService() *AuthenticationSubscriptionService {

	return &AuthenticationSubscriptionService{

	}
}


// AuthenticationSubscriptionSearch returns object by it's UeID
func (s *AuthenticationSubscriptionService) AuthenticationSubscriptionSearch(ueID string) (*models.AuthenticationSubscription, *models.ProblemDetails) {

	return nil, nil
}