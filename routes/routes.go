package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const CHAVE_AUTENTICACAO = "Q0hBVkUgQVVURU5USUNBQ0FP"

type httpHandlerFunc func(http.ResponseWriter, *http.Request)

//AuthApi - realiza a autenticação dos requests
func authAPI(next httpHandlerFunc) httpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("KEY_AUTH")
		if header != CHAVE_AUTENTICACAO {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"errors": "401 - Acesso Negado"}`))
			return
		}
		next(w, r)
	}
}

func routeDefault(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello World!"}`))
}

func GetRoutes() {
	router := mux.NewRouter()
	lsPorta := "9898"

	router.HandleFunc("/", (routeDefault)).Methods("GET")

	fmt.Printf("Aguardando requisições na Porta %s", lsPorta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", lsPorta), router))
}
