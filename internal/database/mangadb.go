package database

import (
	"challenger/internal/entity"
	"database/sql"
)

type MangaDB struct {
	database *sql.DB
}

func Constructor(database *sql.DB) *MangaDB {
	return &MangaDB{database: database}
}

func (mangadb *MangaDB) Create(manga *entity.Manga) (*entity.Manga, error) {
	_, err := mangadb.database.Exec("INSERT INTO mangas (id, name, description) VALUES (?,?,?)", manga.ID, manga.Name, manga.Description)

	if err != nil {
		return nil, err
	}
	return manga, nil
}
