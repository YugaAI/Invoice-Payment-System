package invoice

import (
	"invoice-payment-system/dto"
	"invoice-payment-system/usecase/invoiceUsecase/invoice_read_usecase"
	"invoice-payment-system/usecase/invoiceUsecase/invoice_write_usecase"

	"github.com/gin-gonic/gin"
)

type InvoiceWriteUsecase interface {
	ApproveExecute(cmd invoice_write_usecase.ApproveInvoiceCommand) error
	SubmitExecute(cmd invoice_write_usecase.SubmitInvoiceCommand) error
	CreateExecute(cmd invoice_write_usecase.CreateInvoiceCommand) (uint64, error)
	PaymentExecute(cmd invoice_write_usecase.PayInvoiceCommand) error
}

type InvoiceReadUsecase interface {
	GetListInvoiceExecute(companyID uint64, page, limit int) ([]dto.InvoiceList, error)
	GetInvoiceByIdExecute(q invoice_read_usecase.GetInvoiceDetailQuery) (*dto.InvoiceDetail, error)
	InvoiceDasboardExecute(q invoice_read_usecase.InvoiceDashboardQuery) (*dto.InvoiceDashboard, error)
}

type InvoiceHandler struct {
	*gin.Engine
	UsecaseWrite InvoiceWriteUsecase
	UsecaseRead  InvoiceReadUsecase
}

func NewInvoiceHandler(api *gin.Engine, ucRead InvoiceReadUsecase, ucWrite InvoiceWriteUsecase) *InvoiceHandler {
	return &InvoiceHandler{
		Engine:       api,
		UsecaseRead:  ucRead,
		UsecaseWrite: ucWrite,
	}
}

func (h *InvoiceHandler) RegisterInvoiceRoutes() {
	invoices := h.Group("/invoices")
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
	h.GET("/invoice-dashboard", h.Dashboard)
}
