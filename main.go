package main

import (
	"gin_api_rest/database"
	"gin_api_rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
