package constant

import "time"

const (
	// Claims
	ClaimsUser = "user"

	// JWT
	TokenDefaultDuration        = 1 * time.Hour
	RefreshTokenDefaultDuration = 24 * time.Hour
)
