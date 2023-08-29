Iniciando um novo projeto

mkdir projeto
cd projeto
go mod init projeto


Libs>
go get:
github.com/joho/godotenv

-DRIVERS
go.mongodb.org/mongo-driver/mongo
go.mongodb.org/mongo-driver/mongo/options
gorm.io/driver/mysql
-FS FM
github.com/gin-gonic/gin
-ORM
gorm.io/gorm
-Nodemon like
github.com/githubnemo/CompileDaemon 

cors if not using fullstack way
go get github.com/gin-contrib/cors

CompileDaemon -command="./api"

Estrutura da pastas x packages

P> Controller
    


Nome do arquivo x nome do pacote tem que ser o mesmo

