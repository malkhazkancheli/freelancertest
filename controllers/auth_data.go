package controllers

import (
	"freelancertest/services/auth"

	"github.com/go-openapi/runtime/middleware"

	"freelancertest/restapi/operations/auth_data"
)

// QueryAuthSubsDataHandlerFunc handles GET /subscription-data/{ueId}/authentication-data/authentication-subscription
func QueryAuthSubsDataHandlerFunc(params auth_data.QueryAuthSubsDataParams) middleware.Responder {

	service := auth.NewAuthenticationSubscriptionService()

	payload, problem := service.AuthenticationSubscriptionSearch(params.UeID)
	if problem != nil {
		return auth_data.NewQueryAuthSubsDataDefault(int(problem.Status)).WithPayload(problem)
	}

	return auth_data.NewQueryAuthSubsDataOK().WithPayload(payload)
}