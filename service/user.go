package service

import (
	"github.com/afz204/golang-second-api/dto"
	"github.com/afz204/golang-second-api/entity"
	"github.com/afz204/golang-second-api/repository"
	"github.com/mashingan/smapping"
	"log"
)

//service contract
type UserService interface {
	Update(user dto.UserUpdateDto) entity.User
	Profile(userId string) entity.User
}

type userService struct {
	userRepository repository.UserResponsitory
}

func NewUserService(userRepo repository.UserResponsitory) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDto) entity.User {
	userUpdate := entity.User{}
	err := smapping.FillStruct(&userUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Faild map %v:", err)
	}
	update := service.userRepository.UpdateUser(userUpdate)
	return update
}

func (service *userService) Profile(userId string) entity.User {
	return service.userRepository.ProfileUser(userId)
}
