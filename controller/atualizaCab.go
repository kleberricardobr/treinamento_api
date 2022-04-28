package controller

import (
	"encoding/json"
	"net/http"
	"treinamento-api/dao"
	"treinamento-api/models"
)

func AtualizaCab(w http.ResponseWriter, r *http.Request) {
	var toDo models.ToDoId

	err := json.NewDecoder(r.Body).Decode(&toDo)
	if HasError(w, r, http.StatusBadRequest, "Falha ao recuperar JSon", err) {
		return
	}

	err = dao.AtualizaCab(toDo)
	if HasError(w, r, http.StatusInternalServerError, "Falha ao atualizar cad ToDo", err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"resposta": "ToDo Cadastrado com Sucesso!"}`))
}
