package main

import (
	"fmt"
	"github.com/ismadrade/go-rest-api/database"
	"github.com/ismadrade/go-rest-api/models"
	"github.com/ismadrade/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConecaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
