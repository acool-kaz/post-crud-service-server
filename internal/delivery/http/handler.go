package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/acool-kaz/post-crud-service-server/internal/models"
	"github.com/acool-kaz/post-crud-service-server/internal/service"
)

type Handler struct {
	service *service.Service
}

func InitHandler(service *service.Service) *Handler {
	log.Println("init http handler")
	return &Handler{
		service: service,
	}
}

type status struct {
	Info string `json:"info"`
}

func (h *Handler) InitRoutes() *http.ServeMux {
	log.Println("init routes")
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello, this is root handler")
	})

	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.URL.Path != "/post" {
			h.errorPage(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}

		switch r.Method {
		case http.MethodGet:
			posts, err := h.service.Post.Read(r.Context())
			if err != nil {
				h.errorPage(w, http.StatusInternalServerError, err.Error())
				return
			}

			if err := json.NewEncoder(w).Encode(posts); err != nil {
				h.errorPage(w, http.StatusInternalServerError, err.Error())
				return
			}
		case http.MethodPost:
			var post models.Post
			if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
				h.errorPage(w, http.StatusBadRequest, err.Error())
				return
			}

			if err := h.service.Post.Create(r.Context(), post); err != nil {
				h.errorPage(w, http.StatusInternalServerError, err.Error())
				return
			}

			if err := json.NewEncoder(w).Encode(status{Info: "post created"}); err != nil {
				h.errorPage(w, http.StatusInternalServerError, err.Error())
				return
			}
		default:
			h.errorPage(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		}
	})

	mux.HandleFunc("/post/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(strings.Split(r.URL.Path, "/post/")[1])
		if err != nil {
			h.errorPage(w, http.StatusNotFound, err.Error())
			return
		}

		switch r.Method {
		case http.MethodGet:
			ctx := context.WithValue(r.Context(), models.PostId, id)

			posts, err := h.service.Post.Read(ctx)
			if err != nil {
				h.errorPage(w, http.StatusInternalServerError, err.Error())
				return
			}

			if err := json.NewEncoder(w).Encode(posts); err != nil {
				h.errorPage(w, http.StatusInternalServerError, err.Error())
				return
			}
		case http.MethodDelete:
			if err := h.service.Post.Delete(r.Context(), id); err != nil {
				h.errorPage(w, http.StatusInternalServerError, err.Error())
				return
			}

			if err := json.NewEncoder(w).Encode(status{Info: "post deleted"}); err != nil {
				h.errorPage(w, http.StatusInternalServerError, err.Error())
				return
			}
		case http.MethodPatch:
			var update models.UpdatePost
			if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
				h.errorPage(w, http.StatusBadRequest, err.Error())
				return
			}

			if err := h.service.Post.Update(r.Context(), id, update); err != nil {
				h.errorPage(w, http.StatusInternalServerError, err.Error())
				return
			}

			if err := json.NewEncoder(w).Encode(status{Info: "post updated"}); err != nil {
				h.errorPage(w, http.StatusInternalServerError, err.Error())
				return
			}
		default:
			h.errorPage(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		}
	})

	return mux
}
