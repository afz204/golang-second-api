package controller

import (
	"github.com/afz204/golang-second-api/dto"
	"github.com/afz204/golang-second-api/entity"
	"github.com/afz204/golang-second-api/helper"
	"github.com/afz204/golang-second-api/service"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

//constract AuthController
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

//put service
type authController struct {
	authService service.AuthService
	jwtService service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService: jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDto dto.LoginDto
	if isValidEmail(loginDto.Email) {
		errDto := ctx.ShouldBind(&loginDto)
		if errDto != nil {
			responses := helper.BuildErrorResponse("Failed to process request", errDto.Error(), helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, responses)
			return
		}
		authRes := c.authService.VerifyCredential(loginDto.Email, loginDto.Password)
		if val, ok := authRes.(entity.User); ok {
			generateToken := c.jwtService.GenerateToken(strconv.FormatUint(val.ID, 10))
			val.Token = generateToken
			response := helper.BuildResponse(true, "OK!", val)
			ctx.JSON(http.StatusOK, response)
			return
		}
		response := helper.BuildErrorResponse("Please check credential !", "Invalid Credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	} else {
		response := helper.BuildErrorResponse("Failed to proccess request !", "Invalid Email", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

func isValidEmail (email string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	if !emailRegex.MatchString(email) {
		return false
	}
	parts := strings.Split(email, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDto dto.RegisterDto
	errDto := ctx.ShouldBind(&registerDto)
	if errDto != nil {
		response := helper.BuildErrorResponse("Failed to register !", errDto.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.authService.IsDuplicateEmail(registerDto.Email) {
		response := helper.BuildErrorResponse("Failed to process request!", "Duplicate Email!", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createUser := c.authService.CreateUser(registerDto)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createUser.ID, 10))
		createUser.Token = token
		response := helper.BuildResponse(true, "Created!", createUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
