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

func (mangaService *MangaService) GetById(id string) (*entity.Manga, error) {
	manga, err := mangaService.MangaDB.GetById(id)

	if err != nil {
		return nil, err
	}
	return manga, nil
}

func (mangaService *MangaService) GetAll() ([]*entity.Manga, error) {
	mangas, err := mangaService.MangaDB.GetAll()

	if err != nil {
		return nil, err
	}

	return mangas, err
}

func (mangaService *MangaService) Delete(id string) (result bool) {
	delete := mangaService.MangaDB.Delete(id)

	return delete
}
