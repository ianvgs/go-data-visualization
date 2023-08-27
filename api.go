package main

import (
	"api/initializers"
	"api/routes"
	"fmt"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	routes.HandleRequests()
	fmt.Println("It works")
}
