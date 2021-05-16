package main

import (
	"net/http"

	"github.com/Notarrogantjustbetter/CRUD/v2/database"
	"github.com/Notarrogantjustbetter/CRUD/v2/routes"
	"github.com/Notarrogantjustbetter/CRUD/v2/utils"
)


func main() {
	utils.LoadTemplate()
	database.SetDbEnv()
	database.ConnectDb()
	router := routes.InitRouter()
	http.ListenAndServe(":8080", router)
}
