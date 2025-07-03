package routes

import (
	"github/mrcsmotta1/go-rest-api-alura/controllers"
	"github/mrcsmotta1/go-rest-api-alura/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.GetPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonalidadeByID).Methods("GET")
	r.HandleFunc("/api/personalidades/", controllers.CreatePersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.UpdatePersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", r))
}
