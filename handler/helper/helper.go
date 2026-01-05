package helper

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseID(c *gin.Context) uint64 {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	return id
}

func MustQueryUint(c *gin.Context, key string) uint64 {
	val := c.Query(key)
	if val == "" {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": key + " is required"},
		)
		return 0
	}

	num, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": key + " must be number"},
		)
		return 0
	}

	return num
}

func MustQueryInt(c *gin.Context, key string, def int) int {
	val := c.Query(key)
	if val == "" {
		return def
	}

	num, err := strconv.Atoi(val)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": key + " must be number"},
		)
		return def
	}

	return num
}
