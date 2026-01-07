package invoice

import (
	"invoice-payment-system/usecase/invoice_query"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *InvoiceHandler) Detail(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice id"})
		return
	}

	result, err := h.DetailQuery.Execute(invoice_query.GetInvoiceDetailQuery{
		InvoiceID: id,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *InvoiceHandler) List(c *gin.Context) {
	companyParam := c.Query("company_id")
	if companyParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_id is required"})
		return
	}

	companyID, err := strconv.ParseUint(companyParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid company id"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	result, err := h.ListQuery.Execute(companyID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  result,
		"page":  page,
		"limit": limit,
	})
}

func (h *InvoiceHandler) Dashboard(c *gin.Context) {
	companyParam := c.Query("company_id")
	if companyParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_id is required"})
		return
	}

	companyID, err := strconv.ParseUint(companyParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid company id"})
		return
	}

	result, err := h.DashboardQuery.Execute(invoice_query.InvoiceDashboardQuery{
		CompanyID: companyID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
