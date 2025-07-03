package controllers

import (
	"encoding/json"
	"fmt"
	"github/mrcsmotta1/go-rest-api-alura/database"
	"github/mrcsmotta1/go-rest-api-alura/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

func GetPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetPersonalidadeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)

	http.Error(w, "Personalidade not found", http.StatusNotFound)
}

func CreatePersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	err := json.NewDecoder(r.Body).Decode(&personalidade)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	database.DB.Create(&personalidade)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(personalidade)
}

func DeletePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	if personalidade.ID == 0 {
		http.Error(w, "Personalidade not found", http.StatusNotFound)
		return
	}

	database.DB.Delete(&personalidade)
	w.WriteHeader(http.StatusNoContent)
}
