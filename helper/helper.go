package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseID(c *gin.Context) uint64 {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	return id
}
