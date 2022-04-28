package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"treinamento-api/dao"
	"treinamento-api/models"

	"github.com/gorilla/mux"
)

func GetToDoByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)
	idInt, err := strconv.Atoi(idStr["ID"])
	if HasError(w, r, http.StatusBadRequest, "Falha ao converter parametro ID", err) {
		return
	}

	toDoList, err := dao.GetToDoByID(idInt)

	idErr := http.StatusInternalServerError
	if err == sql.ErrNoRows {
		idErr = http.StatusNotFound
	}

	if HasError(w, r, idErr, "Falha ao recuperar ToDoList", err) {
		return
	}

	ret, err := json.Marshal(toDoList)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ret)

}

func GetAllToDo(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)
	idInt, err := strconv.Atoi(idStr["QTD"])
	if HasError(w, r, http.StatusBadRequest, "Falha ao converter parametro QTD", err) {
		return
	}

	toDoList, err := dao.GetAllToDo(idInt)

	idErr := http.StatusInternalServerError
	if err == sql.ErrNoRows {
		idErr = http.StatusNotFound
	}

	if HasError(w, r, idErr, "Falha ao recuperar ToDoList(All)", err) {
		return
	}

	ret, err := json.Marshal(toDoList)
	if HasError(w, r, http.StatusInternalServerError, "Falha ao converter dados para JSon - ToDoList(All)", err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ret)
}

func CadToDo(w http.ResponseWriter, r *http.Request) {
	var toDo models.ToDoList

	err := json.NewDecoder(r.Body).Decode(&toDo)
	if HasError(w, r, http.StatusBadRequest, "Falha ao recuperar JSon", err) {
		return
	}

	err = dao.CadToDo(toDo)
	if HasError(w, r, http.StatusInternalServerError, "Falha ao cadastrar ToDo", err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"resposta": "Nova lista ToDo cadastrada com sucesso!"}`))
}
