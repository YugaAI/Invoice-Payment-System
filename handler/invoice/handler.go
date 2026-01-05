package invoice

import (
	"invoice-payment-system/usecase/invoice_command"
	"invoice-payment-system/usecase/invoice_query"

	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	CreateCmd  *invoice_command.CreateInvoiceUsecase
	SubmitCmd  *invoice_command.SubmiteInvoiceUsecase
	ApproveCmd *invoice_command.ApproveInvoiceUsecase
	PayCmd     *invoice_command.PayInvoiceUsecase

	DetailQuery    *invoice_query.GetInvoiceDetailUsecase
	ListQuery      *invoice_query.ListInvoiceUsecase
	DashboardQuery *invoice_query.InvoiceDashboardUsecase
}

//func NewInvoiceHandler(create *invoice_command.CreateInvoiceUsecase,
//	submit *invoice_command.SubmiteInvoiceUsecase,
//	approve *invoice_command.ApproveInvoiceUsecase,
//	pay *invoice_command.PayInvoiceUsecase,
//	detail *invoice_query.GetInvoiceDetailUsecase,
//	list *invoice_query.ListInvoiceUsecase,
//	dashboard *invoice_query.InvoiceDashboardUsecase) *InvoiceHandler {
//	return &InvoiceHandler{
//		CreateCmd:      create,
//		SubmitCmd:      submit,
//		ApproveCmd:     approve,
//		PayCmd:         pay,
//		DetailQuery:    detail,
//		ListQuery:      list,
//		DashboardQuery: dashboard,
//	}
//}

func NewRouterInvoice(r *gin.Engine, invoiceHandler *InvoiceHandler) *gin.Engine {
	//r := gin.New()

	// middleware global
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	api := r.Group("/api")
	{
		registerInvoiceRoutes(api, invoiceHandler)
	}

	return r
}

func registerInvoiceRoutes(rg *gin.RouterGroup, h *InvoiceHandler,
) {
	invoices := rg.Group("/invoices")
	{
		// COMMAND
		invoices.POST("/create", h.Create)
		invoices.POST("/:id/submit", h.Submit)
		invoices.POST("/:id/approve", h.Approve)
		invoices.POST("/:id/pay", h.Pay)

		// QUERY
		invoices.GET("/:id", h.Detail)
		invoices.GET("/list", h.List)
	}

	// dashboard biasanya beda concern
	rg.GET("/invoice-dashboard", h.Dashboard)
}
