package main

import (
	"net/http"

	"github.com/thiagohmm/Desafio01FullcycleGoLang/controller"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", controller.PegaDolar)

	http.ListenAndServe(":8080", mux)

}
