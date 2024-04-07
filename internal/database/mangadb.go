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

func (mangadb *MangaDB) GetById(id string) (*entity.Manga, error) {
	var manga entity.Manga

	query := mangadb.database.QueryRow("SELECT Id, Name, Description FROM mangas WHERE ID = ?", id).Scan(&manga.ID, &manga.Name, &manga.Description)

	if query != nil {
		return nil, query
	}

	return &manga, nil
}

func (mangadb *MangaDB) GetAll() ([]*entity.Manga, error) {
	query, err := mangadb.database.Query("SELECT Id, Name, Description FROM mangas")

	if err != nil {
		return nil, err
	}

	defer query.Close()
	var mangas []*entity.Manga

	for query.Next() {
		var manga entity.Manga

		if err := query.Scan(&manga.ID, &manga.Name, &manga.Description); err != nil {
			return nil, err
		}
		mangas = append(mangas, &manga)
	}
	return mangas, nil
}

func (mangadb *MangaDB) Delete(id string) (result bool) {
	delete, err := mangadb.database.Exec("DELETE FROM mangas WHERE ID = ?", id)

	if err != nil {
		return false
	}

	rowsAffect, err := delete.RowsAffected()

	if rowsAffect <= 0 {
		return false
	}

	return true
}

func (mangadb *MangaDB) Update(mangaUpdate *entity.Manga, id string) (*entity.Manga, error) {
	update, err := mangadb.database.Exec("UPDATE mangas SET Name = ?, Description = ? WHERE ID = ?", mangaUpdate.Name, mangaUpdate.Description, id)

	if err != nil {
		return nil, err
	}

	rowsAffect, err := update.RowsAffected()

	if rowsAffect <= 0 {
		return nil, err
	}

	return mangaUpdate, err
}
