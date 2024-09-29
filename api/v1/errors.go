package v1

var (
	// common errors
	ErrSuccess             = newError(0, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")

	// Business Errors
	// Error code convention: 1 + <2-dig module id> + <2-dig error id>

	// User errors
	ErrEmailAlreadyUse    = newError(10001, "The email is already in use.")
	ErrUsernameAlreadyUse = newError(10002, "The username is already in use.")

	// Team errors
	ErrTeamIDAlreadyUse = newError(10101, "The teamID is already in use.")
)
