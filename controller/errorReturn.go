package controller

import (
	"fmt"
	"log"
	"net/http"
)

func HasError(w http.ResponseWriter, r *http.Request, tipoRet int,
	msgErr string, err error) bool {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(tipoRet)
		w.Write([]byte(fmt.Sprintf(`{"falha": "%d - %s"}`, tipoRet, msgErr)))
		log.Printf("Erro no request %v", err)
	}

	return err != nil

}
