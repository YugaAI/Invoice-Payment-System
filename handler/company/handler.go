package company

import (
	"invoice-payment-system/model"

	"github.com/gin-gonic/gin"
)

type CompanyReadUsecase interface {
	GetByIdExecute(id uint64) (*model.Company, error)
	GetListExecute() ([]model.Company, error)
}

type CompanyWriteUsecase interface {
	CreateExecute(name string) (*model.Company, error)
}

type CompanyHandler struct {
	*gin.Engine
	ReadUC  CompanyReadUsecase
	WriteUC CompanyWriteUsecase
}

func NewCompanyHandler(api *gin.Engine, readUC CompanyReadUsecase, writeUC CompanyWriteUsecase) *CompanyHandler {
	return &CompanyHandler{
		Engine:  api,
		ReadUC:  readUC,
		WriteUC: writeUC,
	}
}
func (h *CompanyHandler) RegisterCompanyRoutes() {
	companies := h.Group("/companies")
	{
		companies.POST("/create", h.Create)
		companies.GET("/list", h.List)
		companies.GET("/:id", h.Detail)
	}
}
