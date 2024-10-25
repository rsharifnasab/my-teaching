package repository

import (
	"gorm-postgres/database"
	"gorm-postgres/models"
)

// mockgen: by uber

type Book interface {
	Add(models.Book) (models.Book, error)
	GetAll() ([]models.Book, error)
	Get(title string) (models.Book, error)
	Delete(title string) error
	Update(updated models.Book) (models.Book, error)
}

type gormBook struct {
	DB database.Dbinstance
}

func NewGormBook(db database.Dbinstance) Book {
	return &gormBook{
		DB: db,
	}
}

func (b *gormBook) Add(newBook models.Book) (models.Book, error) {
	result := b.DB.Db.Create(newBook)
	if result.Error != nil {
		return models.Book{}, result.Error
	}

	return newBook, nil
}

func (b *gormBook) GetAll() ([]models.Book, error) {
	var bookList []models.Book
	result := b.DB.Db.Find(&bookList)
	if result.Error != nil {
		return nil, result.Error
	}

	return bookList, nil
}

func (b *gormBook) Get(title string) (models.Book, error) {
	return models.Book{}, nil
}

func (b *gormBook) Delete(title string) error {
	return nil
}

func (b *gormBook) Update(updated models.Book) (models.Book, error) {
	return models.Book{}, nil
}

var _ Book = (*gormBook)(nil)
