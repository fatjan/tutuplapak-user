package handlers

import (
	"github.com/fatjan/tutuplapak/internal/config"
	authRepository "github.com/fatjan/tutuplapak/internal/repositories/auth"
	authUseCase "github.com/fatjan/tutuplapak/internal/usecases/auth"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(cfg *config.Config, db *sqlx.DB, r *gin.Engine) {
	// jwtMiddleware := jwt_helper.JWTMiddleware(cfg.JwtKey)

	r.GET("/", func(c *gin.Context) {
		c.Status(200)
	})

	v1 := r.Group("v1")

	authRepository := authRepository.NewAuthRepository(db)
	authUseCase := authUseCase.NewUseCase(authRepository, cfg)
	authHandler := NewAuthHandler(authUseCase)

	authRouter := v1.Group("")
	authRouter.POST("/register/email", authHandler.RegisterEmail)
	authRouter.POST("/register/phone", authHandler.RegisterPhone)
	authRouter.POST("/login/email", authHandler.LoginEmail)
	authRouter.POST("/login/phone", authHandler.LoginPhone)
}
