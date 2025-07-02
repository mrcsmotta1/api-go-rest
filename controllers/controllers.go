package controllers

import (
	"encoding/json"
	"fmt"
	"github/mrcsmotta1/go-rest-api-alura/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

func GetPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

func GetPersonalidadeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personalidade := range models.Personalidades {
		if fmt.Sprintf("%d", personalidade.ID) == id {
			json.NewEncoder(w).Encode(personalidade)
			return
		}
	}

	http.Error(w, "Personalidade not found", http.StatusNotFound)
}
