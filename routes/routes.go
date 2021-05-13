package routes

import (
	"net/http"

	"github.com/Notarrogantjustbetter/CRUD/v2/database"
	"github.com/Notarrogantjustbetter/CRUD/v2/utils"
	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.ExecuteTemplate(w, "home.html", nil)
	case "POST":
		r.ParseForm()
		name := r.PostForm.Get("Name")
		email := r.PostForm.Get("Email")
		database.CreateUser(name, email)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.ExecuteTemplate(w, "deleteUser.html", nil)
	case "POST":
		r.ParseForm()
		id := r.PostForm.Get("ID")
		database.DeleteUser(id)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.ExecuteTemplate(w, "updateUser.html", nil)
	case "POST":
		r.ParseForm()
		id := r.PostForm.Get("ID")
		key := r.PostForm.Get("Key")
		value := r.PostForm.Get("Value")
		database.UpdateUser(id, key, value)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "users.html", database.GetUsers())
}

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET", "POST")
	router.HandleFunc("/deleteUser", deleteUserHandler).Methods("GET", "POST")
	router.HandleFunc("/updateUser", updateUserHandler).Methods("GET", "POST")
	router.HandleFunc("/getUsers", getUsersHandler).Methods("GET")
	return router
}
