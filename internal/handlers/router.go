package handlers

import (
	"github.com/fatjan/tutuplapak/internal/config"
	"github.com/fatjan/tutuplapak/internal/pkg/jwt_helper"
	authRepository "github.com/fatjan/tutuplapak/internal/repositories/auth"
	userRepository "github.com/fatjan/tutuplapak/internal/repositories/user"
	authUseCase "github.com/fatjan/tutuplapak/internal/usecases/auth"
	userUseCase "github.com/fatjan/tutuplapak/internal/usecases/user"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(cfg *config.Config, db *sqlx.DB, r *gin.Engine) {
	jwtMiddleware := jwt_helper.JWTMiddleware(cfg.JwtKey)

	r.GET("/", func(c *gin.Context) {
		c.Status(200)
	})

	v1 := r.Group("v1")

	authRepository := authRepository.NewAuthRepository(db)
	authUseCase := authUseCase.NewUseCase(authRepository, cfg)
	authHandler := NewAuthHandler(authUseCase)

	userRepository := userRepository.NewUserRepository(db)
	userUseCase := userUseCase.NewUseCase(userRepository)
	userHandler := NewUserHandler(userUseCase)

	authRouter := v1.Group("")
	authRouter.POST("/register/email", authHandler.RegisterEmail)
	authRouter.POST("/register/phone", authHandler.RegisterPhone)
	authRouter.POST("/login/email", authHandler.LoginEmail)
	authRouter.POST("/login/phone", authHandler.LoginPhone)

	userRouter := v1.Group("user")
	userRouter.Use(jwtMiddleware)
	userRouter.GET("/", userHandler.Get)
	userRouter.PATCH("/", userHandler.Update)
}
