package user

import (
	"invoice-payment-system/model"

	"github.com/gin-gonic/gin"
)

type UserReadUsecase interface {
	Login(email, password string) (*model.User, error)
}

type UserWriteUsecase interface {
	SignUp(username, email, password, role string) (*model.User, error)
}

type Handler struct {
	*gin.Engine
	userReadUC  UserReadUsecase
	userWriteUC UserWriteUsecase
}

func NewUserHandler(api *gin.Engine, readUC UserReadUsecase, writeUC UserWriteUsecase) *Handler {
	return &Handler{
		Engine:      api,
		userReadUC:  readUC,
		userWriteUC: writeUC,
	}
}

func (h *Handler) RegisterUserRoutes() {
	users := h.Group("/users")
	{
		users.POST("/signup", h.SignUp)
		users.GET("/login", h.Login)
	}
}
