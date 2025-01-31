package handlers

import (
	"net/http"

	"github.com/fatjan/tutuplapak/internal/dto"
	"github.com/fatjan/tutuplapak/internal/pkg/exceptions"
	"github.com/fatjan/tutuplapak/internal/usecases/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	RegisterEmail(ginCtx *gin.Context)
	RegisterPhone(ginCtx *gin.Context)
	LoginEmail(ginCtx *gin.Context)
	LoginPhone(ginCtx *gin.Context)
}

type authHandler struct {
	authUseCase auth.UseCase
}

func NewAuthHandler(authUseCase auth.UseCase) AuthHandler {
	return &authHandler{authUseCase: authUseCase}
}

func (r *authHandler) RegisterEmail(ginCtx *gin.Context) {
	var authRequest dto.AuthRequest
	if err := ginCtx.BindJSON(&authRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	authResponse, err := r.authUseCase.Register(ginCtx.Request.Context(), &authRequest, false)
	if err != nil {
		ginCtx.JSON(exceptions.MapToHttpStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusCreated, &dto.AuthResponseEmail{
		Email: authResponse.Email,
		Token: authResponse.Token,
	})
}

func (r *authHandler) RegisterPhone(ginCtx *gin.Context) {
	var authRequest dto.AuthRequest
	if err := ginCtx.BindJSON(&authRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	authResponse, err := r.authUseCase.Register(ginCtx.Request.Context(), &authRequest, true)
	if err != nil {
		ginCtx.JSON(exceptions.MapToHttpStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusCreated, &dto.AuthResponsePhone{
		Phone: authResponse.Phone,
		Token: authResponse.Token,
	})
}

func (r *authHandler) LoginEmail(ginCtx *gin.Context) {
	var authRequest dto.AuthRequest
	if err := ginCtx.BindJSON(&authRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	authResponse, err := r.authUseCase.Login(ginCtx.Request.Context(), &authRequest, false)
	if err != nil {
		ginCtx.JSON(exceptions.MapToHttpStatusCode(err), gin.H{"error": err.Error()})
		return
	}
	ginCtx.JSON(http.StatusOK, &dto.AuthResponseEmail{
		Email: authResponse.Email,
		Token: authResponse.Token,
	})
}

func (r *authHandler) LoginPhone(ginCtx *gin.Context) {
	var authRequest dto.AuthRequest
	if err := ginCtx.BindJSON(&authRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	authResponse, err := r.authUseCase.Login(ginCtx.Request.Context(), &authRequest, true)
	if err != nil {
		ginCtx.JSON(exceptions.MapToHttpStatusCode(err), gin.H{"error": err.Error()})
		return
	}
	ginCtx.JSON(http.StatusOK, &dto.AuthResponsePhone{
		Phone: authResponse.Phone,
		Token: authResponse.Token,
	})
}
