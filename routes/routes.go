package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"treinamento-api/controller"
	"treinamento-api/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func routeDefault(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"resposta": "Conectado!"}`))
}

func GetListenPort() (ret string) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Falha ao carregar as váriaveis de conexão %v", err)
	}

	return os.Getenv("LISTEN_PORT")
}

func GetRoutes() {
	router := mux.NewRouter()
	lsPorta := GetListenPort()

	router.HandleFunc("/", middleware.ValidateRequestKey(http.HandlerFunc(routeDefault))).Methods("GET")

	router.HandleFunc("/GET_TO_DO/{ID}",
		middleware.ValidateRequestKey(http.HandlerFunc(controller.GetToDoByID))).Methods("GET")

	router.HandleFunc("/GET_TO_DO_ALL/{QTD}",
		middleware.ValidateRequestKey(http.HandlerFunc(controller.GetAllToDo))).Methods("GET")

	router.HandleFunc("/CAD_TO_DO",
		middleware.ValidateRequestKey(http.HandlerFunc(controller.CadToDo))).Methods("POST")

	router.HandleFunc("/DATA_HORA/{LOCAL1}/{LOCAL2}",
		middleware.ValidateRequestKey(http.HandlerFunc(controller.GetDataHora))).Methods("GET")

	router.HandleFunc("/TIME_ZONES",
		middleware.ValidateRequestKey(http.HandlerFunc(controller.GetTimeZones))).Methods("GET")

	router.HandleFunc("/ATU_CAB",
		middleware.ValidateRequestKey(http.HandlerFunc(controller.AtualizaCab))).Methods("PUT")

	log.Printf("--->>>Aguardando requisições na Porta %s<<<---", lsPorta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", lsPorta), router))
}
