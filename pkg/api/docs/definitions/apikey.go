package apidocs

import (
	"github.com/grafana/grafana/pkg/api/dtos"
	"github.com/grafana/grafana/pkg/models"
)

// swagger:route GET /auth/keys apikeys getAPIkeys
//
// Get auth keys
//
// Will return auth keys.
//
// Responses:
// 200: apikeyResponse
// 401: notAuthorizedError
// 403: forbiddenError
// 404: notFoundError
// 500: internalServerError

// swagger:route POST /auth/keys apikeys addAPIkey
//
// Creates an API key
//
// Will return details of the created API key
//
// Responses:
// 200: newAPIkeyResponse
// 400: badRequestError
// 401: notAuthorizedError
// 403: forbiddenError
// 403: quotaReachedError
// 409: conflictError
// 500: internalServerError

// swagger:route DELETE /auth/keys/{id} apikeys deleteAPIkey
//
// 409: conflictError
// Delete API key
//
// Responses:
// 200: okResponse
// 409: conflictError
// 401: notAuthorizedError
// 403: forbiddenError
// 404: notFoundError
// 500: internalServerError

// swagger:parameters getAPIkeys
type getAPIkeysParams struct {
	// Show expired keys
	// in:query
	// required:false
	// default:false
	IncludeExpired bool `json:"includeExpired"`
}

// swagger:model
type apikeyResponse []*models.ApiKeyDTO

// swagger:parameters addAPIkey
type addAPIkeyParams struct {
	// in:body
	Body addAPIkeyCmd
}

// swagger:model
type addAPIkeyCmd models.AddApiKeyCommand

// swagger:model
type newAPIkeyResponse dtos.NewApiKeyResult

// swagger:parameters deleteAPIkey
type deleteAPIkeyParams struct {
	// in:path
	// required:true
	ID int64 `json:"id"`
}
