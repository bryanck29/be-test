package request

import "github.com/google/uuid"

type PostInsertUser struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required,oneof=user admin"`
}

type GetUser struct {
	UserId uuid.UUID `param:"user_id" validate:"required"`
}

type DeleteUser struct {
	UserId uuid.UUID `param:"user_id" validate:"required"`
}

type PutUser struct {
	UserId   uuid.UUID `param:"user_id" validate:"required"`
	Name     string    `json:"name,omitempty" validate:"omitempty"`
	Username string    `json:"username,omitempty" validate:"omitempty"`
	Password string    `json:"password,omitempty" validate:"omitempty"`
	Role     string    `json:"role,omitempty" validate:"omitempty,eq=user|eq=admin"`
}
