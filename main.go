package main

import (
	"fmt"
	"github/mrcsmotta1/go-rest-api-alura/database"
	"github/mrcsmotta1/go-rest-api-alura/models"
	"github/mrcsmotta1/go-rest-api-alura/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{ID: 1, Nome: "Albert Einstein", Historia: "Físico teórico, desenvolvedor da teoria da relatividade."},
		{ID: 2, Nome: "Marie Curie", Historia: "Pioneira na pesquisa sobre radioatividade, primeira mulher a ganhar um Prêmio Nobel."},
		{ID: 3, Nome: "Nelson Mandela", Historia: "Líder anti-apartheid, primeiro presidente negro da África do Sul."},
	}

	database.Connect()

	fmt.Println("Starting server on port 8080...")
	routes.HandleRequests()
}
