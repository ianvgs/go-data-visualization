package examplepkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SayHi() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"Result": "SayHi",
		})

	}
}
