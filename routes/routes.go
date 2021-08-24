package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Deikioveca/CRUD/v2/database"
	"github.com/Deikioveca/CRUD/v2/utils"
	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.ExecuteTemplate(w, "home.html", nil)
	case "POST":
		r.ParseForm()
		user := database.User{}
		user.Name = r.PostForm.Get("Name")
		err := database.PostgresDb{}.CreateUser(&user)
		if err != nil {
			fmt.Fprint(w, "Failed to create user")
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.ExecuteTemplate(w, "deleteUser.html", nil)
	case "POST":
		r.ParseForm()
		user := database.User{}
		id := r.PostForm.Get("ID")
		strToInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Fprint(w, "Failed to convert string to int")
		}
		user.ID = uint(strToInt)
		err = database.PostgresDb{}.DeleteUser(&user)
		if err != nil {
			fmt.Fprint(w, "Failed to delete user")
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.ExecuteTemplate(w, "updateUser.html", nil)
	case "POST":
		r.ParseForm()
		user := database.User{}
		id := r.PostForm.Get("id")
		user.Name = r.PostForm.Get("newName")
		strToInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Fprint(w, "Failed to convert id")
		}
		user.ID = uint(strToInt)
		err = database.PostgresDb{}.UpdateUser(&user)
		if err != nil {
			fmt.Fprint(w, "Failed to update user")
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	context, err := database.PostgresDb{}.GetAllUsers()
	if err != nil {
		fmt.Fprint(w, "Failed to get users")
	}
	utils.ExecuteTemplate(w, "users.html", context)
}

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET", "POST")
	router.HandleFunc("/deleteUser", deleteUserHandler).Methods("GET", "POST")
	router.HandleFunc("/updateUser", updateUserHandler).Methods("GET", "POST")
	router.HandleFunc("/getUsers", getUsersHandler).Methods("GET")
	return router
}
