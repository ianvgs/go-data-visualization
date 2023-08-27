package sentiment_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
)

func AnalyseTextSentimeter() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Declaro estrutura
		var body struct {
			TextToAnalysis string
		}
		//Faço o Bind
		c.BindJSON(&body)
		//Lanço o erro se a propriedade nao existir
		if body.TextToAnalysis == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"warning": "Nao foi encontrado os parâmetros necessários para realizar essa tarefa.",
			})

		}

		//Parseio o text
		parsedText := sentitext.Parse(body.TextToAnalysis, lexicon.DefaultLexicon)
		//Analiso a popularidade
		result := sentitext.PolarityScore(parsedText)

		/* fmt.Println("Positive:", result.Positive)
		fmt.Println("Negative:", result.Negative)
		fmt.Println("Neutral:", result.Neutral) */

		c.JSON(http.StatusOK, gin.H{
			"Negative:": result.Negative,
			"Positive:": result.Positive,
			"Neutral:":  result.Neutral,
		})

	}
}
