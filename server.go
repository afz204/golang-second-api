package main

import (
	"github.com/afz204/golang-second-api/config"
	"github.com/afz204/golang-second-api/controller"
	"github.com/afz204/golang-second-api/repository"
	"github.com/afz204/golang-second-api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabase()
	userReponsitory repository.UserResponsitory = repository.NewUserReponsitory(db)
	jwtService service.JWTService = service.NewJWTService()
	authService service.AuthService = service.NewAuthService(userReponsitory)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main()  {
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
