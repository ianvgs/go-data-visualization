package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func envMySqlString() string {

	varMode := os.Getenv("GO_ENV")
	log.Println("MODE:", varMode)

	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file at envMySqlString. DEV_MODE")
		}

	}

	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	log.Println("Username:", DB_USERNAME)

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	return dsn
}

func ConnectToDB() {

	var err error
	dsn := envMySqlString()

	//log.Println("Dsn:", dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//Desse jeito ele não deixa a aplicação rodar
	/* if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	} */

	if err != nil {
		log.Println("Erro ao conectar ao banco de dados:", err)
		// handle the error here, for example:
		// return an error message to the user
		// retry the connection after a certain amount of time
		/* By using log.Println instead of log.Fatal, the program will continue running even if the database connection fails. You can then handle the error in a way that makes sense for your application, such as showing an error message to the user or retrying the connection after a certain amount of time. */
	}

	/* DB.AutoMigrate(&models.Post{}) */

}
