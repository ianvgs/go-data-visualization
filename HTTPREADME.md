Status:
http.StatusBadRequest
http.StatusUnauthorized
http.StatusMovedPermanently
http.StatusOK
http.StatusInternalServerError
http.StatusNotFound


PORT=3000
GO_ENV=dev

API_URL="localhost:8080"

#database config
DB_USERNAME = "ian"
DB_PASSWORD = "Genuine1#"
DB_NAME = "go_analytics"
DB_HOST = "44.212.43.124"
DB_PORT = "3306"


MONGODB_URI= 'mongodb+srv://f4324234:170892@mdbcluster.tgepnhk.mongodb.net/nest-api?retryWrites=true&w=majority'


returns types of c.Context
JSON


fullstack:
HTML
c.HTML(http.StatusBadRequest, "404.html", gin.H{"content": "Page not found."})

c.Redirect(http.StatusMovedPermanently, "/login")
c.Abort()

