package controller

import (
	"challenger/internal/entity"
	"challenger/internal/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type MangaController struct {
	MangaService *services.MangaService
}

func Constructor(mangaService *services.MangaService) *MangaController {
	return &MangaController{MangaService: mangaService}
}

func (mangaHandle *MangaController) Create(writer http.ResponseWriter, request *http.Request) {
	var manga entity.Manga
	err := json.NewDecoder(request.Body).Decode(&manga)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := mangaHandle.MangaService.CreateService(manga.Name, manga.Description)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(result)
	writer.WriteHeader(http.StatusCreated)
}

func (mangaHandle *MangaController) GetById(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		http.Error(writer,
			"Id is required", http.StatusBadRequest)
		return
	}

	manga, err := mangaHandle.MangaService.GetById(id)

	if err != nil {
		switch err.Error() {
		case "sql: no rows in result set":
			http.Error(writer, "Manga not found", http.StatusNotFound)
			return

		default:
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	json.NewEncoder(writer).Encode(manga)
}

func (mangaHandle *MangaController) GetAll(writer http.ResponseWriter, request *http.Request) {
	mangas, err := mangaHandle.MangaService.GetAll()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if mangas == nil {
		http.Error(writer, "Mangas not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(writer).Encode(mangas)
}
func (mangaHandle *MangaController) Delete(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		http.Error(writer,
			"Id is required", http.StatusBadRequest)
		return
	}

	delete := mangaHandle.MangaService.Delete(id)

	if delete == false {
		http.Error(writer, "Manga not found to delete", http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

func (mangaHandle *MangaController) Update(writer http.ResponseWriter, request *http.Request) {
	var manga entity.Manga
	err := json.NewDecoder(request.Body).Decode(&manga)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := mangaHandle.MangaService.UpdateService(manga.ID, manga.Name, manga.Description)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(result)
}
