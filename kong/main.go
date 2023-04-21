package main

import (
	"fmt"
	"net/http"
	"github.com/kevinbtian/interview/kong/handler"
	"github.com/kevinbtian/interview/kong/database"
)

func main() {
    cfg := database.Config{
        ProjectID:  "my-project",
        InstanceID: "my-instance",
        Database:   "my-database",
        Username:   "my-username",
        Password:   "my-password",
    }

    db := database.NewDatabase(cfg)
    defer db.Close()

	h := handler.NewHandler(db)
	http.HandleFunc("/service", h.GetServicesHandler)

	fmt.Println("Starting server on hocalhost:8080...")
	http.ListenAndServe(":8080", nil)
}
