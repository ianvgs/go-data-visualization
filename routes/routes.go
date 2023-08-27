package routes

import (
	data_controller "api/controllers/data_controller"
	sentiment_controller "api/controllers/sentiment_controller"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	public := r.Group("/")
	publicRoutes(public)

	r.Run(":8080")
}

func publicRoutes(g *gin.RouterGroup) {

	g.GET("/sentiment", sentiment_controller.AnalyseTextSentimeter())
	g.GET("/csv", data_controller.ReadCSV())

	g.GET("/plotter/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		filePath := filepath.Join("output", filename)
		fmt.Println(filePath)

		// Check if the file exists
		_, err := os.Stat(filePath)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "File not found",
			})
			return
		}

		// Set the content type to image/png
		c.Header("Content-Type", "image/png")

		// Serve the file as-is
		c.File(filePath)
	})

	g.POST("/csver", data_controller.GenerateChartsFromGivenCsvAndTargetColumn())

}
