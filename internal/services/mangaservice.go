package services

import (
	"challenger/internal/database"
	"challenger/internal/entity"
)

type MangaService struct {
	MangaDB database.MangaDB
}

func Constructor(mangaDB database.MangaDB) *MangaService {
	return &MangaService{MangaDB: mangaDB}
}

func (mangaService *MangaService) CreateService(name, description string) (*entity.Manga, error) {
	manga := entity.Constructor(name, description)

	_, err := mangaService.MangaDB.Create(manga)

	if err != nil {
		return nil, err
	}

	return manga, nil
}
