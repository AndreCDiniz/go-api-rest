package main

import (
	"fmt"

	"github.com/AndreCDiniz/api-go-rest/database"
	"github.com/AndreCDiniz/api-go-rest/models"
	"github.com/AndreCDiniz/api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "hasjasahdjkaj"},
		{Id: 2, Nome: "Nome 2", Historia: "asjldalksjaksdj"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
