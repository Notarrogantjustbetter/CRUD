package main

import (
	"net/http"

	"github.com/Notarrogantjustbetter/CRUD/v2/database"
	"github.com/Notarrogantjustbetter/CRUD/v2/server"
	"github.com/Notarrogantjustbetter/CRUD/v2/utils"
)


func main() {
	utils.LoadTemplate()
	database.InitDb()
	router := server.InitServer()
	http.ListenAndServe(":8080", router)
}