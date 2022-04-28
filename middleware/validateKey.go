package middleware

import (
	"log"
	"net/http"
)

const CHAVE_AUTENTICACAO = "Q0hBVkUgQVVURU5USUNBQ0FP"

type httpHandlerFunc func(http.ResponseWriter, *http.Request)

/*
   Fonte:
   https://stackoverflow.com/questions/27234861/correct-way-of-getting-clients-ip-addresses-from-http-request
*/

func readUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func writeUnauthorized(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(`{"falha": "401 - Acesso Negado"}`))
	log.Printf("Tentativa de acesso nao autorizado! IP: %s", readUserIP(r))
}

func ValidateRequestKey(next http.Handler) httpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("KEY_AUTH") != CHAVE_AUTENTICACAO {
			writeUnauthorized(w, r)
			return
		}

		next.ServeHTTP(w, r)
	}
}
