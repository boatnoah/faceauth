package main

import "net/http"

func (a *app) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All operations are OK\n"))
}
