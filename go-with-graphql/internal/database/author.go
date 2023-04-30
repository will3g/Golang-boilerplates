package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Author struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	Email       string
}

func AuthorInstance(db *sql.DB) *Author {
	return &Author{db: db}
}

func (authorPointer *Author) Create(name, description, email string) (Author, error) {
	id := uuid.New().String()
	myQuery := "INSERT INTO authors (id, name, description, email) VALUES ($1, $2, $3, $4)"
	_, err := authorPointer.db.Exec(myQuery, id, name, description, email)

	if err != nil {
		return Author{}, err
	}

	return Author{
		ID:          id,
		Name:        name,
		Description: description,
		Email:       email,
	}, nil
}

func (authorPointer *Author) FindAll() ([]Author, error) {
	myQuery := "SELECT id, name, description, email FROM authors"
	rows, err := authorPointer.db.Query(myQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	authors := []Author{}

	for rows.Next() {
		var id, name, description, email string
		if err := rows.Scan(&id, &name, &description, &email); err != nil {
			return nil, err
		}
		authors = append(authors, Author{
			ID:          id,
			Name:        name,
			Description: description,
			Email:       email,
		})
	}

	return authors, nil
}

func (authorPointer *Author) FindByArticleID(ArticleID string) (Author, error) {
	var id, name, description, email string
	myQuery := "SELECT authors.id, authors.name, authors.description, authors.email FROM authors INNER JOIN articles ON authors.id = articles.author_id WHERE articles.id = $1"
	if err := authorPointer.db.QueryRow(myQuery, ArticleID).Scan(&id, &name, &description, &email); err != nil {
		return Author{}, err
	}

	return Author{
		ID:          id,
		Name:        name,
		Description: description,
		Email:       email,
	}, nil
}
