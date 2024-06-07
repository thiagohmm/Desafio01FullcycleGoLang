package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/thiagohmm/Desafio01FullcycleGoLang/service"
)

func PegaDolar(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	moeda := service.Usdbrl{}
	cotacao, error := moeda.GetUsdbrl(ctx)
	if error != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("Context deadline exceeded: %v\n", error)

			w.WriteHeader(http.StatusRequestTimeout)
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)

}
