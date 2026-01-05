package invoice

import (
	"invoice-payment-system/handler/helper"
	"invoice-payment-system/usecase/invoice_query"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *InvoiceHandler) Detail(c *gin.Context) {
	id := helper.ParseID(c)

	res, err := h.DetailQuery.Execute(invoice_query.GetInvoiceDetailQuery{
		InvoiceID: id,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *InvoiceHandler) List(c *gin.Context) {
	companyID := helper.MustQueryUint(c, "company_id")

	status := c.Query("status")
	var statusPtr *string
	if status != "" {
		statusPtr = &status
	}

	res, err := h.ListQuery.Execute(invoice_query.ListInvoiceQuery{
		CompanyID: companyID,
		Status:    statusPtr,
		Limit:     helper.MustQueryInt(c, "limit", 10),
		Offset:    helper.MustQueryInt(c, "offset", 0),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *InvoiceHandler) Dashboard(c *gin.Context) {
	companyID := helper.MustQueryUint(c, "company_id")

	res, err := h.DashboardQuery.Execute(invoice_query.InvoiceDashboardQuery{
		CompanyID: companyID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
