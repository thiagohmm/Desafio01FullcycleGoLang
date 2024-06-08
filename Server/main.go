package main

import (
	"net/http"

	"github.com/thiagohmm/Desafio01FullcycleGoLang/controller"
	"github.com/thiagohmm/Desafio01FullcycleGoLang/db"
)

func main() {
	db.InitDB("./usdbrl.db")
	mux := http.NewServeMux()

	mux.HandleFunc("GET /cotacao/", controller.PegaDolar)
	mux.HandleFunc("GET /cotacao", controller.PegaDolar)
	http.ListenAndServe(":8080", mux)

}
