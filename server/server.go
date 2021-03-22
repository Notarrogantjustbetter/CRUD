package server

import (
	"net/http"

	"github.com/Notarrogantjustbetter/CRUD/v2/database"
	"github.com/Notarrogantjustbetter/CRUD/v2/utils"
	"github.com/gorilla/mux"
)


type Server struct {
	router *mux.Router
}

func InitServer() *mux.Router {
	s := &Server{
		router: mux.NewRouter(),
	}
	s.routes()
	return s.router
}

func (s Server) routes() {
	s.router.HandleFunc("/", homeHandler().ServeHTTP).Methods("GET", "POST")
	s.router.HandleFunc("/users", getUsersHandler().ServeHTTP).Methods("GET")
	s.router.HandleFunc("/updateUser", updateUserHandler().ServeHTTP).Methods("GET", "POST")
	s.router.HandleFunc("/deleteUser", deleteUserHandler().ServeHTTP).Methods("GET", "POST")
}

func homeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.PostForm.Get("Name")
		email := r.PostForm.Get("Email")
		database.User{}.CreateUser(name, email)
		utils.ExecuteTemplate(w, "home.html", nil)
	}
}

func getUsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ExecuteTemplate(w, "users.html", database.User{}.GetUsers())
	}
}

func updateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		id := r.PostForm.Get("ID")
		key := r.PostForm.Get("Key")
		value := r.PostForm.Get("Value")
		database.User{}.UpdateUser(id, key, value)
		utils.ExecuteTemplate(w, "updateUser.html", nil)
	}
}

func deleteUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		id := r.PostForm.Get("ID")
		database.User{}.DeleteUser(id)
		utils.ExecuteTemplate(w, "deleteUser.html", nil)
	}
}
