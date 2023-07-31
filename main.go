package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "historia 2"},
	}

	database.ConectaDB()

	fmt.Println("Iniciando o servior REST com go")
	routes.HandleRequest()
}
