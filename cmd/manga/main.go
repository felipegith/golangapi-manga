package main

import (
	"challenger/internal/controller"
	"challenger/internal/database"
	"challenger/internal/services"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	db, err := sql.Open("mysql", "")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	mangaDB := database.Constructor(db)
	mangaService := services.Constructor(*mangaDB)
	mangaController := controller.Constructor(mangaService)

	router := chi.NewRouter()
	router.Post("/create", mangaController.Create)

	http.ListenAndServe(":8080", router)
}
