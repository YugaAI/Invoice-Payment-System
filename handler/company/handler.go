package company

import (
	"invoice-payment-system/usecase/company_command"
	"invoice-payment-system/usecase/company_query"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	CreateCmd *company_command.CreateCompanyUsecase

	GetQuery  *company_query.GetCompanyUsecase
	ListQuery *company_query.ListCompanyUsecase
}

func NewRouterCompany(r *gin.Engine, CompanyHandler *CompanyHandler) *gin.Engine {
	//r := gin.New()

	// middleware global
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	api := r.Group("/api")
	{
		registerCompanyRoutes(api, CompanyHandler)
	}

	return r
}
func registerCompanyRoutes(rg *gin.RouterGroup, h *CompanyHandler) {
	companies := rg.Group("/companies")
	{
		companies.POST("/create", h.Create)
		companies.GET("/list", h.List)
		companies.GET("/:id", h.Detail)
	}
}
