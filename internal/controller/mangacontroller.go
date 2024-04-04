package controller

import (
	"challenger/internal/entity"
	"challenger/internal/services"
	"encoding/json"
	"net/http"
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
}
