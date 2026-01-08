package invoice

import (
	"fmt"
	"invoice-payment-system/helper"
	"invoice-payment-system/usecase/invoiceUsecase/invoice_write_usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createInvoiceReq struct {
	CompanyID uint64                                    `json:"company_id" binding:"required"`
	Items     []invoice_write_usecase.CreateInvoiceItem `json:"items" binding:"required"`
}

func (h *InvoiceHandler) Create(c *gin.Context) {
	var req createInvoiceReq
	fmt.Printf("%+v\n", req.Items)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.UsecaseWrite.CreateExecute(invoice_write_usecase.CreateInvoiceCommand{
		CompanyID: req.CompanyID,
		Items:     req.Items,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"invoice_id": id})
}

func (h *InvoiceHandler) Submit(c *gin.Context) {
	id := helper.ParseID(c)

	err := h.UsecaseWrite.SubmitExecute(invoice_write_usecase.SubmitInvoiceCommand{
		InvoiceID: id,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

type approveReq struct {
	Approver string `json:"approver"`
}

func (h *InvoiceHandler) Approve(c *gin.Context) {
	id := helper.ParseID(c)

	var req approveReq
	_ = c.ShouldBindJSON(&req)

	err := h.UsecaseWrite.ApproveExecute(invoice_write_usecase.ApproveInvoiceCommand{
		InvoiceID: id,
		Approver:  req.Approver,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

type payReq struct {
	Method string `json:"method" binding:"required"`
	RefNo  string `json:"ref_no" binding:"required"`
}

func (h *InvoiceHandler) Pay(c *gin.Context) {
	id := helper.ParseID(c)

	var req payReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.UsecaseWrite.PaymentExecute(invoice_write_usecase.PayInvoiceCommand{
		InvoiceID: id,
		Method:    req.Method,
		RefNo:     req.RefNo,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
