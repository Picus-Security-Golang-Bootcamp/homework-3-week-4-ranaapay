package storage

import (
	"PicusHomework3/src/storage/helper"
	"PicusHomework3/src/storage/storageType"
	"errors"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (a *AuthorRepository) GetById(id int) (*storageType.Author, error) {
	var author storageType.Author
	result := a.db.First(&author, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &author, nil
}

func (a *AuthorRepository) FindByName(name string) []storageType.Author {
	var authors []storageType.Author
	a.db.Where("firstName LIKE ? ", "%"+name+"%").
		Find(&authors)
	return authors
}

func (a *AuthorRepository) GetAllAuthorsWithBookInfo() (storageType.Author, error) {
	var authors []storageType.Author
	result := a.db.Preload("books").Find(&authors)
	if result.Error != nil {
		return storageType.Author{}, result.Error
	}
	return authors[1], nil
}

func (a *AuthorRepository) InsertSampleData() error {
	//csvFile := os.Getenv("AUTHOR_CSV")
	authors, err := helper.ReadCSVForAuthor("author.csv")
	if err != nil {
		return err
	}
	for _, author := range authors {
		a.db.Where(storageType.Author{Name: author.Name, LastName: author.LastName}).
			Create(&author)
	}
	return nil
}

func (a *AuthorRepository) Migrations() {
	a.db.AutoMigrate(&storageType.Author{})
}
