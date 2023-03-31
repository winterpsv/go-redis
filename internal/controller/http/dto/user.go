package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateUserDTO struct {
	Nickname  string `json:"nickname" validate:"required,min=3,max=20,alphanum"`
	FirstName string `json:"first_name" validate:"required,alphaunicode,max=50"`
	LastName  string `json:"last_name" validate:"required,alphaunicode,max=50"`
	Password  string `json:"password" validate:"required,min=8,max=25"`
	Role      string `json:"role" validate:"required,oneof=user moderator admin"`
}

type UpdateUserDTO struct {
	FirstName string `json:"first_name" validate:"required,alphaunicode,max=50"`
	LastName  string `json:"last_name" validate:"required,alphaunicode,max=50"`
	Role      string `json:"role" validate:"required,oneof=user moderator admin"`
}

type UpdateUserPasswordDTO struct {
	Password string `json:"password" validate:"required,min=8,max=25"`
}

type CreateTokenDTO struct {
	Nickname string `json:"nickname" validate:"required,min=3,max=20,alphanum"`
	Password string `json:"password" validate:"required,min=8,max=25"`
}

type VoteUserDTO struct {
	Value int `json:"value" validate:"required,oneof=-1 1"`
}

type UserDTO struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Nickname     string             `json:"nickname"`
	FirstName    string             `json:"first_name"`
	LastName     string             `json:"last_name"`
	PasswordHash string             `json:"password"`
	CreatedAt    int64              `json:"created_at"`
	UpdatedAt    int64              `json:"updated_at"`
	DeletedAt    int64              `json:"deleted_at"`
	Role         string             `json:"role"`
	Active       bool               `json:"active"`
	Votes        int                `json:"votes"`
}
