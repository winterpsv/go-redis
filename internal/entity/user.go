package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Nickname     string             `json:"nickname" bson:"nickname"`
	FirstName    string             `json:"first_name" bson:"first_name"`
	LastName     string             `json:"last_name" bson:"last_name"`
	PasswordHash string             `json:"password" bson:"password"`
	CreatedAt    int64              `json:"created_at" bson:"created_at"`
	UpdatedAt    int64              `json:"updated_at" bson:"updated_at"`
	DeletedAt    int64              `json:"deleted_at" bson:"deleted_at"`
	Role         string             `json:"role" bson:"role"`
	Active       bool               `json:"active" bson:"active"`
	Votes        []UserVote         `json:"votes" bson:"votes"`
}

// UserVote model
type UserVote struct {
	VoterID   primitive.ObjectID `json:"voter_id" bson:"voter_id"`
	VoteValue int                `json:"vote_value" bson:"vote_value"`
	VotedAt   int64              `json:"voted_at" bson:"voted_at"`
}
