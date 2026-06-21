package services

import (
	"backend-go-mysql/dto/request"
	"backend-go-mysql/dto/response"
	"backend-go-mysql/models"
	"backend-go-mysql/repositories"
	"backend-go-mysql/utils"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return UserService{UserRepo: userRepo}
}

func (s *UserService) CreateUser(userDto request.UserDto) (response.UserResponse, error) {

	// ✅ generate UUID
	userUUID := uuid.New().String()

	// ✅ hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.UserResponse{}, err
	}

	user := models.User{
		UserId:       userUUID,
		Email:        userDto.Email,
		FirstName:    userDto.FirstName,
		MidName:      utils.StringPtr(userDto.MidName),
		LastName:     userDto.LastName,
		PhoneNo:      userDto.PhoneNo,
		Password:     string(hashedPassword),
		ImageProfile: utils.StringPtr(userDto.ImageProfile),
		CreatedAt:    utils.NowTH(),
	}

	result, err := s.UserRepo.CreateUser(user)
	if err != nil {
		return response.UserResponse{}, errors.New(err.Error())
	}

	// ✅ map response (ไม่มี password)
	userResponse := response.UserResponse{
		UserId:       result.UserId,
		Email:        result.Email,
		FirstName:    result.FirstName,
		MidName:      result.MidName,
		LastName:     result.LastName,
		PhoneNo:      result.PhoneNo,
		ImageProfile: result.ImageProfile,
	}

	return userResponse, nil
}

func (s *UserService) GetProfileByUserId(userId string) (response.UserResponse, error) {
	user, err := s.UserRepo.FindByUserId(userId)

	if err != nil {
		return response.UserResponse{}, errors.New("user not found")
	}

	userResponse := response.UserResponse{
		UserId:       user.UserId,
		Email:        user.Email,
		FirstName:    user.FirstName,
		MidName:      user.MidName,
		LastName:     user.LastName,
		PhoneNo:      user.PhoneNo,
		ImageProfile: user.ImageProfile,
	}

	return userResponse, nil
}
