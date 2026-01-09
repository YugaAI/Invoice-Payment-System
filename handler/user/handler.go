package user

import (
	"invoice-payment-system/auth"
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
	//pasetoSvc   *auth.PasetoService
	pasetoSvc *auth.PasetoPublicService
}

func NewUserHandler(api *gin.Engine, readUC UserReadUsecase, writeUC UserWriteUsecase, pasetoSvc *auth.PasetoPublicService) *Handler {
	return &Handler{
		Engine:      api,
		userReadUC:  readUC,
		userWriteUC: writeUC,
		pasetoSvc:   pasetoSvc,
	}
}

func (h *Handler) RegisterUserRoutes() {
	users := h.Group("/users")
	{
		users.POST("/signup", h.SignUp)
		users.GET("/login", h.Login)
	}
}
