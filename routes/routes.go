package routes

import (
	examplepkg "api/controllers"
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
	// Example
	g.GET("/sayhi", examplepkg.SayHi())
	g.GET("/saybye", examplepkg.SayBye())
	// End Example

	/* {
		"textToAnalysis":"loved it"
	} */
	g.GET("/sentiment", sentiment_controller.AnalyseTextSentimeter())

	//Examples for functions
	g.GET("/csv", data_controller.ReadCSV())

	/* {
		"colTarget":"ColunaIdade",
		"csv": arquivo.csv
	} */
	//Enviar arquivo csv e a coluna que deve ser usada pra gerar as visualizações
	g.POST("/csver", data_controller.GenerateChartsFromGivenCsvAndTargetColumn())

	//Método pra recuperar os gráficos pela url
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

}
