package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ramiro/API/entities"
	"log"
	"net/http"
)

type IResp struct {
	Error   bool
	Message string
}

func Create(w http.ResponseWriter, r *http.Request) {
	var todo entities.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println("Erro ao fazer decode do json %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) // 3 retornos responseWriter onde ele vai escrever a msg, a menssagem de erro em si e o utimo o status code
		return
	}

	id, err := entities.Insert(todo)

	var resp map[string]IResp

	if err != nil {
		resp = map[string]IResp{
			"Error":   {Error: true},
			"Message": {Message: fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err)},
		}
	} else {
		resp = map[string]IResp{
			"Error":   {Error: false},
			"Message": {Message: fmt.Sprintf("Todo inserido com sucesso %d", id)},
		}
	}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
