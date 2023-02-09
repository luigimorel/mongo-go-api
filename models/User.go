package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignUpInput struct {
	Username        string             `json:"user_name" bson:"username" binding:"required"`
	FirstName       string             `json:"first_name" bson:"first_name" binding:"required"`
	LastName        string             `json:"last_name" bson:"last_name" binding:"required"`
	Email           string             `json:"email" bson:"email" binding:"required"`
	Bio             string             `json:"bio" bson:"bio" binding:"required"`
	DateOfBirth     primitive.DateTime `json:"date_of_birth" bson:"date_of_birth" binding:"required"`
	Password        string             `json:"password" bson:"password" binding:"required,min=8"`
	PasswordConfirm string             `json:"passwordConfirm" bson:"passwordConfirm,omitempty" binding:"required"`
	Verified        bool               `json:"verified" bson:"verified"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
}

type SignInInput struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type DBResponse struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Username        string             `json:"user_name" bson:"username"`
	FirstName       string             `json:"first_name" bson:"first_name"`
	LastName        string             `json:"last_name" bson:"last_name"`
	Email           string             `json:"email" bson:"email"`
	Bio             string             `json:"bio" bson:"bio"`
	Password        string             `json:"password" bson:"password"`
	DateOfBirth     time.Time          `json:"date_of_birth" bson:"date_of_birth"`
	PasswordConfirm string             `json:"passwordConfirm,omitempty" bson:"passwordConfirm,omitempty"`
	Verified        bool               `json:"verified" bson:"verified"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
}

type UserResponse struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username    string             `json:"user_name,omitempty" bson:"username,omitempty"`
	FirstName   string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName    string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Bio         string             `json:"bio,omitempty" bson:"bio,omitempty"`
	DateOfBirth time.Time          `json:"date_of_birthomitempty" bson:"date_of_birth,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

func FilteredResponse(user *DBResponse) UserResponse {
	user.DateOfBirth, _ = time.Parse("2006-01-02", user.DateOfBirth.String())
	return UserResponse{
		ID:          user.ID,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Bio:         user.Bio,
		DateOfBirth: user.DateOfBirth,
		Username:    user.Username,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}
