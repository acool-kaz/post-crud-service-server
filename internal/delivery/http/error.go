package http

import (
	"log"
	"net/http"
)

func (h *Handler) errorPage(w http.ResponseWriter, code int, err string) {
	log.Println(err)
	http.Error(w, http.StatusText(code), code)
}
