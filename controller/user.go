package controller

import (
	"fmt"
	"github.com/afz204/golang-second-api/dto"
	"github.com/afz204/golang-second-api/helper"
	"github.com/afz204/golang-second-api/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController interface {
	Update(ctx *gin.Context)
	Profile(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (ct *userController) Update(context *gin.Context) {
	var userUpdateDto dto.UserUpdateDto
	errDto := context.ShouldBind(&userUpdateDto)
	if errDto != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDto.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	getToken := context.GetHeader("Authorization")
	token, errToken := ct.jwtService.ValidateToken(getToken)
	if errToken != nil {
		panic(errToken.Error())
	}
	claimToken := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claimToken["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}

	userUpdateDto.ID = id
	user := ct.userService.Update(userUpdateDto)
	res := helper.BuildResponse(true, "OK!", user)
	context.JSON(http.StatusOK, res)
}
func (ct *userController) Profile(context *gin.Context) {
	getToken := context.GetHeader("Authorization")
	token, err := ct.jwtService.ValidateToken(getToken)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := ct.userService.Profile(id)
	res := helper.BuildResponse(true, "OK!", user)
	context.JSON(http.StatusOK, res)
}
