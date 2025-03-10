package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"AzureGoMySQLNebular/pkg/auth"
	"AzureGoMySQLNebular/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	r.Use(auth.JwtAuthentication)
	routes.RegisterArticleRoutes(r)
	routes.RegisterStatusRoutes(r)
	routes.RegisterUserRoutes(r)
	routes.RegisterArticleRoutes(r)
	routes.RegisterSubscribersRoutes(r)
	routes.RegisterAdministratorRoutes(r)
	http.Handle("/", r)
	handler := cors.Default().Handler(r)
	fmt.Println("Starting web server...")
	log.Fatal(http.ListenAndServe(os.Getenv("GOLANG_PORT"), handler))
}
