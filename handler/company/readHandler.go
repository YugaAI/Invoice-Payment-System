package company

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *CompanyHandler) Detail(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	company, err := h.ReadUC.GetByIdExecute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "companyUsecase not found"})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (h *CompanyHandler) List(c *gin.Context) {
	companies, err := h.ReadUC.GetListExecute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, companies)
}
