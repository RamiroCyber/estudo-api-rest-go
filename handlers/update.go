package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/ramiro/API/entities"
	"log"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("Erro ao fazer parse do if: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) // 3 retornos responseWriter onde ele vai escrever a msg, a menssagem de erro em si e o utimo o status code
		return
	}

	var todo entities.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println("Erro ao fazer decode do json %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) // 3 retornos responseWriter onde ele vai escrever a msg, a menssagem de erro em si e o utimo o status code
		return
	}
	rows, err := entities.Update(int64(id), todo)
	if err != nil {
		log.Println("Erro ao fazer decode do json %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if rows > 1 {
		log.Println("Error: foram atualizados %d registros", rows)
	}
	resp := map[string]IResp{
		"Error":   {Error: false},
		"Message": {Message: "Dados atualizados com sucesso!"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
