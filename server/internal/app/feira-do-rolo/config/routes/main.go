package routes

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func configureAllRoutes() *mux.Router {
	r := mux.NewRouter()
	OfertasRouting(r)
	return r
}

func configureCORSAndServe(r *mux.Router) {
	allAllowedMethods := make([]string, 0)
	allAllowedMethods = append(allAllowedMethods, "GET", "POST", "PUT", "PATCH", "DELETE")
	r.Methods("GET", "POST", "PUT", "PATCH", "DELETE")

	methods := handlers.AllowedMethods(allAllowedMethods)
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-request-with", "Access-Control-Allow-Origin", "Content-Type"})

	srv := &http.Server{
		Handler: handlers.CORS(methods, origins, headers)(r),
		Addr:    ":8080",
	}
	log.Fatal(srv.ListenAndServe())
}

func ConfigureServer() {
	r := configureAllRoutes()
	configureCORSAndServe(r)
}
