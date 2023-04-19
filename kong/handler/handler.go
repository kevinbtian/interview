package handler

import (
	// "database/sql"
	"fmt"
	"net/http"
	"strconv"
)

type Handler struct {
	// db *sql.DB
}

func NewHandler() *Handler {
	return &Handler{}
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
		output := fmt.Sprintf("q: %s, size: %v, page: %v", q, size, page)
		fmt.Fprintf(w, output)

	default:
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}
}
