package storage

import (
	"PicusHomework3/src/storage/helper"
	"PicusHomework3/src/storage/storageType"
	"errors"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (b *BookRepository) GetById(id int) (*storageType.Book, error) {
	var book storageType.Book
	result := b.db.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &book, nil
}

func (b *BookRepository) FindByName(name string) []storageType.Book {
	var books []storageType.Book
	b.db.Where("Name LIKE ? ", "%"+name+"%").
		Find(&books)
	return books
}

func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&storageType.Book{})
}

func (b *BookRepository) InsertSampleData() error {
	//jsonFile := os.Getenv("BOOK_JSON")
	books, err := helper.ReadJSONForBook("book.json")
	if err != nil {
		return err
	}
	b.db.Create(&books)
	return nil
}
