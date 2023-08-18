package model

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID `json:"id" bson:"id,omitempty"`
	Name      string    `json:"name" bson:"name,omitempty"`
	Username  string    `json:"username" bson:"username,omitempty"`
	Password  string    `json:"-" bson:"password,omitempty"`
	Role      string    `json:"role" bson:"role,omitempty"`
	CreatedAt int64     `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt int64     `json:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt int64     `json:"deleted_at" bson:"deleted_at,omitempty"`
}
