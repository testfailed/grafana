package apidocs

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// The response message
	// in: body
	Body ErrorResponseBody `json:"body"`
}

type ErrorResponseBody struct {
	// a human readable version of the error
	// required: true
	Message string `json:"message"`

	// Error An optional detailed description of the actual error. Only included if running in developer mode.
	Error string `json:"error"`

	// Status An optional status to denote the cause of the error.
	//
	// For example, a 412 Precondition Failed error may include additional information of why that error happened.
	Status string `json:"status"`
}
type SuccessResponseBody struct {
	Message string `json:"message,omitempty"`
}

// OKResponse
//
// swagger:response okResponse
type OKResponse struct {
	// in: body
	Body SuccessResponseBody `json:"body"`
}

// An UnauthorizedError is an error that is used when the request is not authenticated.
//
// swagger:response unauthorisedError
type UnauthorizedError GenericError

// A ForbiddenError is used if the user/token has insufficient permission to access the requested resource.
//
// swagger:response forbiddenError
type ForbiddenError GenericError

// A NotFoundError is used when the requested resource was not found.
//
// swagger:response notFoundError
type NotFoundError GenericError

// A BadRequestError is used when the request is invalid and it cann't be processed.
//
// swagger:response badRequestError
type BadRequestError GenericError

// ConflictError
//
// swagger:response conflictError
type ConflictError GenericError

// PreconditionFailedError
//
// swagger:response preconditionFailedError
type PreconditionFailedError GenericError

// UnprocessableEntityError
//
// swagger:response unprocessableEntityError
type UnprocessableEntityError GenericError

// InternalServerError
//
// swagger:response internalServerError
type InternalServerError GenericError

// QuotaReachedError Quota reached for this resource.
//
// swagger:response quotaReachedError
type QuotaReachedError GenericError

// NotAuthorizedError The request is not authenticated.
//
// swagger:response notAuthorizedError
type NotAuthorizedError GenericError

/*

// Create/update dashboard response.
// swagger:model PostDashboardResponse
type PostDashboardResponse struct {
	// Status status of the response.
	// required: true
	// example: success
	Status string `json:"status"`

	// Slug The slug of the dashboard.
	// required: true
	// example: my-dashboard
	Slug string `json:"title"`

	// Version The version of the dashboard.
	// required: true
	// example: 2
	Verion int64 `json:"version"`

	// UID The unique identifier (uid) of the created/updated dashboard.
	// required: true
	// example: nHz3SXiiz
	UID string `json:"uid"`

	// URL The relative URL for accessing the created/updated dashboard.
	// required: true
	// example: /d/nHz3SXiiz/my-dashboard
	URL string `json:"url"`
}

// Get home dashboard response.
// swagger:model GetHomeDashboardResponse
type GetHomeDashboardResponse struct {
	// swagger:allOf
	// required: false
	DashboardFullWithMeta

	// swagger:allOf
	// required: false
	DashboardRedirect
}
*/
