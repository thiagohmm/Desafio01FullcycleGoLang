package controller

import "net/http"

func PegaDolar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Word!"))
}
