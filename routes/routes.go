package routes

import (
	"fmt"
	"log"
	"net/http"
	"treinamento-api/controller"
	"treinamento-api/middleware"

	"github.com/gorilla/mux"
)

func routeDefault(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"resposta": "Conectado!"}`))
}

func GetRoutes() {
	router := mux.NewRouter()
	lsPorta := "9898"

	router.HandleFunc("/", middleware.ValidateRequestKey(http.HandlerFunc(routeDefault))).Methods("GET")

	router.HandleFunc("/GET_TO_DO/{ID}",
		middleware.ValidateRequestKey(http.HandlerFunc(controller.GetToDoByID))).Methods("GET")

	router.HandleFunc("/GET_TO_DO_ALL/{QTD}",
		middleware.ValidateRequestKey(http.HandlerFunc(controller.GetAllToDo))).Methods("GET")

	log.Printf("--->>>Aguardando requisições na Porta %s<<<---", lsPorta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", lsPorta), router))
}
