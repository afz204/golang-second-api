package service

import (
	"github.com/afz204/golang-second-api/dto"
	"github.com/afz204/golang-second-api/entity"
	"github.com/afz204/golang-second-api/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

//is a contract
type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDto) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserResponsitory
}

func NewAuthService(userRepo repository.UserResponsitory) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}

func (service *authService) VerifyCredential(email string, pass string) interface{}  {
	res := service.userRepository.VerifyCredential(email, pass)
	if val, ok := res.(entity.User); ok {
		comparePass := comparePass(val.Password, []byte(pass))
		if val.Email == email && comparePass {
			return res
		}
		return false
	}
	return false
}
func (service *authService) CreateUser(user dto.RegisterDto) entity.User  {
	userCreate := entity.User{}
	mapping := smapping.FillStruct(&userCreate, smapping.MapFields(&user))
	if mapping != nil {
		panic("Failed to create user!")
	}
	res := service.userRepository.Insertuser(userCreate)
	return res
}
func (service *authService) FindByEmail(email string) entity.User  {
	return service.userRepository.FindByEmail(email)
}
func (service *authService) IsDuplicateEmail(email string) bool  {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePass(hashedPwd string, originPass []byte) bool  {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, originPass)
	if err != nil {
		return false
	}
	return true
}