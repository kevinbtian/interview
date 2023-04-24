package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kevinbtian/interview/kong/handler"
	"github.com/kevinbtian/interview/kong/database"
)

func main() {
    cfg := database.Config{
        Instance: os.Getenv("CLOUD_SQL_CONNECTION_NAME"),
        Database: os.Getenv("DB_NAME"),
        Username: os.Getenv("CLOUD_SQL_USER"),
        Password: os.Getenv("CLOUD_SQL_PASSWORD"),
    }

    db := database.NewDatabase(cfg)
    defer db.Close()

	h := handler.NewHandler(db)
	http.HandleFunc("/service", h.GetServicesHandler)
	http.HandleFunc("/create", h.CreateServicesHandler)

	fmt.Println("Starting server on hocalhost:8080...")
	http.ListenAndServe(":8080", nil)
}
