package main

import "net/http"

func AccountServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("0"))
}
