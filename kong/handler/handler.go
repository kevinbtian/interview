package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kevinbtian/interview/kong/database"

)

type Handler struct {
	db *database.Database
}

func NewHandler(db *database.Database) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetServicesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		q := r.URL.Query().Get("q")
		size, err := strconv.Atoi(r.URL.Query().Get("size"))
		if err != nil {
			http.Error(w, "invalid 'size' number", http.StatusBadRequest)
			return
		}
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			http.Error(w, "invalid 'page' number", http.StatusBadRequest)
			return
		}

		// Call database stuff here.
		// output := fmt.Sprintf("q: %s, size: %v, page: %v", q, size, page)
		output, err := h.db.GetServices(q, size, page)
		if err != nil {
			http.Error(w, "could not query services from database", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, output)

	default:
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) CreateServicesHandler(w http.ResponseWriter, r *http.Request) {
	err := h.db.CreateServices(); if err != nil {
		output := fmt.Sprintf("could not create services table: %v", err)
		http.Error(w, output, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "successfully created services table.")
	return
}