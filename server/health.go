package server

import "net/http"

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("In am all good yay!!"))
}
