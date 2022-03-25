package storageType

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name     string `json:"firstName"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Books    []Book `gorm:"foreignKey:AuthorId"`
}
