package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const END_POINT = "http://worldtimeapi.org/api/timezone/"

type response struct {
	DateTime string `json:"datetime"`
	TimeZone string `json:"timezone"`
}

func GetDataHora(w http.ResponseWriter, r *http.Request) {
	local1 := mux.Vars(r)["LOCAL1"]
	local2 := mux.Vars(r)["LOCAL2"]
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s", END_POINT, local1, local2))

	if HasError(w, r, resp.StatusCode, fmt.Sprintf("Falha ao consultar %s/%s/%s", END_POINT, local1, local2), err) {
		return
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if HasError(w, r, resp.StatusCode, "Falha ao recuperar dados retornados pela consulta de API externa", err) {
		return
	}

	var responseJson response
	json.Unmarshal(responseData, &responseJson)

	log.Printf(`Data e Hora Recuperados %s`, responseJson.DateTime)
	log.Printf(`TimeZone Recuperado %s`, responseJson.TimeZone)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

func GetTimeZones(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(END_POINT)

	if HasError(w, r, resp.StatusCode, fmt.Sprintf("Falha ao consultar %s", END_POINT), err) {
		return
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if HasError(w, r, resp.StatusCode, "Falha ao recuperar dados retornados pela consulta de API externa", err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
