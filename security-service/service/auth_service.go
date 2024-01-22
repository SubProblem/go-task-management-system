package service

import (
	"errors"
	"fmt"
	db "subproblem/security-service/database"
	"subproblem/security-service/dto"
	"subproblem/security-service/jwt"
	"subproblem/security-service/util"
)

type AuthService struct {
	db db.PostgresDb
}

func NewAuthService(database *db.PostgresDb) *AuthService {
	return &AuthService{
		db: *database,
	}
}

func (auth *AuthService) Login(userReqeust dto.LoginRequestDto) (string, error) {

	user, err := auth.db.FindUserByEmail(userReqeust.Email)
	fmt.Printf("user: %v\n", user)
	if err != nil {
		return "", err
	}

	if !util.ComparePassowrd(user.Password, userReqeust.Password) || user.Email != userReqeust.Email {
		return "", errors.New("Bad Credentials")
	}

	token, err := jwt.GenerateToken(userReqeust.Email, user.ID)

	if err != nil {
		return "", nil
	}

	return token, nil
}

func (auth *AuthService) Register(userRequest dto.UserRequestDto) error {

	user, err := auth.db.FindUserByEmail(userRequest.Email)

	if err != nil {
		return err
	}
	fmt.Printf("user: %v\n", user)

	if user != nil {
		return errors.New("Registration Failed")
	}

	newUser := dto.ToUser(userRequest)

	hashedPassword, err := util.HashPassword(newUser.Password)
	if err != nil {
		return err
	}

	newUser.Password = hashedPassword

	if err := auth.db.AddUser(newUser); err != nil {
		return errors.New("Registration Failed, try again")
	}

	return nil
}
