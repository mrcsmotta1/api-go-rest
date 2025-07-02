package main

import (
	"fmt"
	"github/mrcsmotta1/go-rest-api-alura/routes"
)

func main() {
	fmt.Println("Starting server on port 8080...")
	routes.HandleRequests()
}
