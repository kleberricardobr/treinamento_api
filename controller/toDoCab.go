package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"treinamento-api/dao"

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ret)

}
