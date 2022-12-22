package main

import (
	"backend/database"
	"backend/pkg/mysql"
	"backend/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mysql.Databaseinit()
	database.RunMigration()
	r := mux.NewRouter()
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())
	fmt.Println("Server running on port 5000")
	http.ListenAndServe("localhost: 5000", r)
}
