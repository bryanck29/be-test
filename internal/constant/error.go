package constant

import (
	"fmt"
	"net/http"
)

var (
	ErrParsingRequest         = fmt.Errorf("failed to parse request")
	ErrNoAuth                 = fmt.Errorf("no authorization")
	ErrInvalidAuth            = fmt.Errorf("invalid authorization")
	ErrInvalidSession         = fmt.Errorf("invalid session")
	ErrDataNotFound           = fmt.Errorf("data not found")
	ErrDataExists             = fmt.Errorf("data already exists")
	ErrCreatingUser           = fmt.Errorf("error while trying to create new user")
	ErrGettingUser            = fmt.Errorf("error while fetching user data")
	ErrDeletingUser           = fmt.Errorf("error while deleting user data")
	ErrInvalidLoginCredential = fmt.Errorf("invalid login credential")
	ErrInvalidAuthToken       = fmt.Errorf("invalid auth token")
	ErrInvalidAccess          = fmt.Errorf("access is not permitted for your role")
	ErrLogin                  = fmt.Errorf("error while trying to sign in")
	ErrRefreshLogin           = fmt.Errorf("error while trying to refresh auth session")
)

// CommonErrorMap maps errors
var CommonErrorMap = map[error]int{
	ErrParsingRequest:         http.StatusUnprocessableEntity,
	ErrNoAuth:                 http.StatusBadRequest,
	ErrDataNotFound:           http.StatusNotFound,
	ErrDataExists:             http.StatusConflict,
	ErrInvalidAuth:            http.StatusUnauthorized,
	ErrInvalidSession:         http.StatusUnauthorized,
	ErrInvalidLoginCredential: http.StatusBadRequest,
	ErrInvalidAuthToken:       http.StatusUnauthorized,
	ErrInvalidAccess:          http.StatusUnauthorized,
}
