Status:
http.StatusBadRequest
http.StatusUnauthorized
http.StatusMovedPermanently
http.StatusOK
http.StatusInternalServerError
http.StatusNotFound


returns types of c.Context
JSON


fullstack:
HTML
c.HTML(http.StatusBadRequest, "404.html", gin.H{"content": "Page not found."})

c.Redirect(http.StatusMovedPermanently, "/login")
c.Abort()

