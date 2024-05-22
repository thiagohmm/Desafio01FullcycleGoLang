package main

import (
	"net/http"

	"github.com/thiagohmm/Desafio01FullcycleGoLang.git/controller"
)

func main() {
	http.HandleFunc("/", controller.PegaDolar)
	http.ListenAndServe(":8080", nil)
}
