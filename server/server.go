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

func (s Server) InitServer() *mux.Router {
	myServer := &Server{
		router: mux.NewRouter(),
	}
	myServer.routes()
	return myServer.router
}

func (s Server) routes() {
	s.router.HandleFunc("/", s.homeHandler().ServeHTTP).Methods("GET", "POST")
	s.router.HandleFunc("/deleteUser", s.deleteUserHandler().ServeHTTP).Methods("GET", "POST")
	s.router.HandleFunc("/updateUser", s.updateUserHandler().ServeHTTP).Methods("GET", "POST")
	s.router.HandleFunc("/getUsers", s.getUsersHandler().ServeHTTP).Methods("GET")
}

func (s Server) homeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			utils.ExecuteTemplate(w, "home.html", nil)
		case "POST":
			r.ParseForm()
			name := r.PostForm.Get("Name")
			email := r.PostForm.Get("Email")
			database.PostgresDB{}.CreateUser(name, email)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func (s Server) deleteUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			utils.ExecuteTemplate(w, "deleteUser.html", nil)
		case "POST":
			r.ParseForm()
			id := r.PostForm.Get("ID")
			database.PostgresDB{}.DeleteUser(id)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func (s Server) updateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			utils.ExecuteTemplate(w, "updateUser.html", nil)
		case "POST":
			r.ParseForm()
			id := r.PostForm.Get("ID")
			key := r.PostForm.Get("Key")
			value := r.PostForm.Get("Value")
			database.PostgresDB{}.UpdateUser(id, key, value)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func (s Server) getUsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ExecuteTemplate(w, "users.html", database.PostgresDB{}.GetUsers())
	}
}
