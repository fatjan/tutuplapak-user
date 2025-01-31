package handlers

import (
	"log"
	"net/http"

	"github.com/fatjan/tutuplapak/internal/dto"
	"github.com/fatjan/tutuplapak/internal/pkg/exceptions"
	"github.com/fatjan/tutuplapak/internal/usecases/user"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Get(ginCtx *gin.Context)
	Update(ginCtx *gin.Context)
}

type userHandler struct {
	userUseCase user.UseCase
}

func (r *userHandler) Get(ginCtx *gin.Context) {
	userId, exists := ginCtx.Get("user_id")
	if !exists {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user id"})
		return
	}

	id := userId.(int)
	userRequest := dto.UserRequest{
		UserID: id,
	}

	userResponse, err := r.userUseCase.GetUser(ginCtx.Request.Context(), &userRequest)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, userResponse)
}

func (r *userHandler) Update(ginCtx *gin.Context) {
	if ginCtx.ContentType() != "application/json" {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content type"})
		return
	}

	var userRequest dto.UserPatchRequest

	userId, exists := ginCtx.Get("user_id")
	if !exists {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user id"})
		return
	}

	userIDInt := userId.(int)

	if err := ginCtx.ShouldBindJSON(&userRequest); err != nil {
		log.Println(err.Error())
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userResponse, err := r.userUseCase.UpdateUser(ginCtx.Request.Context(), userIDInt, &userRequest)
	if err != nil {
		ginCtx.JSON(exceptions.MapToHttpStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, userResponse)
}

func NewUserHandler(userUseCase user.UseCase) UserHandler {
	return &userHandler{userUseCase: userUseCase}
}
