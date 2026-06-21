package services

import (
	"backend-go-mysql/dto/response"
	"backend-go-mysql/repositories"
	"backend-go-mysql/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo   repositories.UserRepository
	JwtService *utils.JwtService
}

func NewAuthService(
	userRepo repositories.UserRepository,
	jwtService *utils.JwtService,
) *AuthService {
	return &AuthService{
		UserRepo:   userRepo,
		JwtService: jwtService,
	}
}

func (s *AuthService) Login(email, password string) (response.UserResponse, error) {
	user, err := s.UserRepo.FindByEmail(email)

	if err != nil {
		return response.UserResponse{}, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return response.UserResponse{}, errors.New("password incorrect")
	}

	token, err := s.JwtService.GenerateToken(user.UserId)
	if err != nil {
		return response.UserResponse{}, err
	}

	// ✅ map ตรงนี้
	userResponse := response.UserResponse{
		UserId:       user.UserId,
		Email:        user.Email,
		FirstName:    user.FirstName,
		MidName:      user.MidName,
		LastName:     user.LastName,
		PhoneNo:      user.PhoneNo,
		Token:        token,
		ImageProfile: user.ImageProfile,
	}

	return userResponse, nil
}
