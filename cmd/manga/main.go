package main

import (
	"challenger/internal/controller"
	"challenger/internal/database"
	"challenger/internal/services"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	godotenv.Load()
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	mangaDB := database.Constructor(db)
	mangaService := services.Constructor(*mangaDB)
	mangaController := controller.Constructor(mangaService)

	router := chi.NewRouter()

	router.Put("/manga", mangaController.Update)
	router.Get("/mangas", mangaController.GetAll)
	router.Post("/create", mangaController.Create)
	router.Get("/manga/{id}", mangaController.GetById)
	router.Delete("/manga/{id}", mangaController.Delete)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
