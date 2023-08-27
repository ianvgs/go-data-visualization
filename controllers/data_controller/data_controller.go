package data_controller

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var outputFolder = "./output"

func ReadCSV() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Declaro estrutura
		/* var body struct {
			TextToAnalysis string
		} */
		//Faço o Bind
		/* 	c.BindJSON(&body) */
		//Lanço o erro se a propriedade nao existir
		/* if body.TextToAnalysis == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"warning": "Nao foi encontrado os parâmetros necessários para realizar essa tarefa.",
			})

		} */

		//Open CSV
		csvFile, err := os.Open("data/inventory.csv")

		fmt.Println("Passei aqui")

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Passei aqui 2")

		defer csvFile.Close()
		//Read CSV
		df := dataframe.ReadCSV(csvFile)
		/* fmt.Println(df) */
		fmt.Println("Passei aqui 3")
		//Shape of dataset
		/* row, col := df.Dims() */
		/* 	fmt.Println("Shape of df, rows:", row, "columns:", col) */

		/* Tamanho das linhas) */
		/* 	fmt.Println("Rows:", df.Nrow()) */

		/* Tamanho das colunas*/
		/* fmt.Println("Cols:", df.Ncol()) */

		/*Nomes das colunas */
		/* 		fmt.Println("Nomes:", df.Names()) */

		/*Tipos das colunas */
		/* fmt.Println("Tipos:", df.Types()) */

		//Criar dataframe
		/* 	fmt.Println("Describe:", df.Describe()) */

		//Selecionar coluna por nome (case sensitive)
		/* 	fmt.Println("SELECT:", df.Select("Model")) */

		//Selecionar coluna por index
		/* 	fmt.Println("SELECT:", df.Select(0)) */

		//Selecionar varias colunas
		/* var colunas = []string{"Model", "Make"} */

		/* var colunas = []string
		colunas = append(colunas, "Model")
		colunas = append(colunas, "Make") */

		/* 		fmt.Println("SELECT:", df.Select(colunas)) */

		//Selecionar Rows por index
		/* 	fmt.Println("SELECT ROWS:", df.Subset(0)) */

		//Aplicando FUNÇÕES
		//ds = data series

		ds := df.Col("Make")
		fmt.Println(ds)
		fmt.Printf("%T \n", ds)

		//APPLY FUNCTION MEAN

		//Gera o valor
		/* 	ds := df.Col("Year") */
		/* 		fmt.Println(ds)
		   		fmt.Printf("%T \n", ds) */
		//Gera a media
		/* 		dsMean := ds.Mean()
		   		fmt.Println(dsMean)

		   		fmt.Println(ds.IsNaN()) */

		//using stats for mean
		statsMean := stat.Mean(ds.Float(), nil)
		fmt.Println(statsMean)

		// Use the Filter method
		filteredDF := df.Filter(dataframe.F{
			Colname:    "Make",
			Comparator: "==",
			Comparando: "Ford",
		})

		fmt.Println(filteredDF)

		c.JSON(http.StatusOK, gin.H{
			"Result": "è nois que ta",
		})

	}

}

/*
var r regression.regression
r.SetObserved("rating")
r.SetVar(0, "Sugars")

//Loop csv records adding the training data

fo i, record := range trainingData {

	//skip the header

	if i == 0 {
		continue
	}

	ratingVal, err := strconv.ParseFloat(record[0],64)
	if err != nil {
		log.Fatal(err)
	}
	sugarsVal, err := strconv.ParseFloat(record[2],64)
	if err != nil {
		log.Fatal(err)
	}

	//Add these point to the regression value
	r.Train(regression.DataPoint(ratingVal, []float64{sugarsVal}))



} */

func GenerateChartsFromGivenCsvAndTargetColumn() gin.HandlerFunc {

	return func(c *gin.Context) {
		/* $EXEMPLO ABRIR CSV LOCAL */
		//Open csv
		/* csvFile, err := os.Open("data/float.csv")
		if err != nil {
			log.Fatal(err)
		} */
		//Remember to close file when returning this function
		/* defer csvFile.Close() */

		/* $$$ 1) Recebe o csv, a coluna target, abre e lê o arquivo */

		//Obtem o valor colTarget do formulário enviado
		receivedColTarget := c.PostForm("colTarget")
		if receivedColTarget == "" {
			c.JSON(400, gin.H{"error": "The property colTarget was not informed in form."})
			return
		}

		//Procura pelo csv chamado csv no form
		csvFile, err := c.FormFile("csv")
		if err != nil {
			c.JSON(400, gin.H{"error": "CSV file not provided"})
			return
		}

		// Open the uploaded file
		srcFile, err := csvFile.Open()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to open uploaded file"})
			return
		}
		defer srcFile.Close()

		//Le no dataframe o csv
		df := dataframe.ReadCSV(srcFile)

		/* $$$ 2) Formata as entradas necessárias para gerar o gráfico */
		//Define a labal para os dados/coluna do eixo Y
		yColumnLabel := receivedColTarget
		//Assimilate Y Val as selected column above
		yVals := df.Col(yColumnLabel).Float()

		// Create an empty array of strings to save path to graphs
		var graphUrls []string

		for _, colName := range df.Names() {

			//Skip the same used Y Column
			if colName == yColumnLabel {
				continue
			}

			pts := make(plotter.XYs, df.Nrow())

			/* for i, floatVal := range df.Col(colName).Float() {
				pts[i].X = floatVal
				pts[i].Y = yVals[i]
			}
			*/

			//Verifica se a coluna que vai ser analisada é numérica, se não for, skipa ela
			for i, floatVal := range df.Col(colName).Float() {
				if !math.IsNaN(floatVal) && !math.IsInf(floatVal, 0) {
					pts[i].X = floatVal
					pts[i].Y = yVals[i]
				} else {
					fmt.Printf("Error at index %d: Invalid floating-point value\n", i)
					continue
				}
			}

			p := plot.New()

			//Label X
			p.X.Label.Text = colName

			//LabelY
			p.Y.Label.Text = yColumnLabel

			p.Add(plotter.NewGrid())

			s, err := plotter.NewScatter(pts)
			if err != nil {
				log.Fatal("NewScatter Error")
			}

			s.GlyphStyle.Color = color.RGBA{R: 233, B: 0, A: 255}
			s.GlyphStyle.Radius = vg.Points(3)

			p.Add(s)

			newFileName := receivedColTarget + "versus" + colName + "_scatter.png"

			err = p.Save(4*vg.Inch, 4*vg.Inch, filepath.Join(outputFolder, newFileName))

			endpointString := os.Getenv("API_URL") + "/plotter/" + newFileName
			graphUrls = append(graphUrls, endpointString)

			if err != nil {
				log.Fatal(err)
			}

		}

		c.JSON(http.StatusOK, gin.H{
			"links": graphUrls,
		})

	}

}
