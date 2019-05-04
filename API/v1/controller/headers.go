package controller

import "net/http"

// SetHeaderJSON define o header do parâmetro 'w http.ResponseWriter' como 'application/json'
func SetHeaderJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}
