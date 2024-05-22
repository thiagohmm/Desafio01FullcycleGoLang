package controller

import (
	"encoding/json"
	"net/http"

	"github.com/thiagohmm/Desafio01FullcycleGoLang/service"
)

func PegaDolar(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	moeda := service.Usdbrl{}
	cotacao, error := moeda.GetUsdbrl()
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)
}
