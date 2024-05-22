package main

import (
	"net/http"

	"github.com/thiagohmm/Desafio01FullcycleGoLang/controller"
)

func main() {
	http.HandleFunc("/", controller.PegaDolar)
	http.ListenAndServe(":8080", nil)
}
