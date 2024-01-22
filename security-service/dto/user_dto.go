package dto

import (
	"subproblem/security-service/models"
)

type UserResponseDto struct {
	ID        int    `json:"id"`
	Firstname string `json:"fisrtname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type UserRequestDto struct {
	Firstname string `json:"fisrtname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToUserDto(user models.User) *UserResponseDto {

	userDto := &UserResponseDto{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}

	return userDto
}

func ToUser(user UserRequestDto) *models.User {

	newUser := &models.User{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
	}

	return newUser
}
