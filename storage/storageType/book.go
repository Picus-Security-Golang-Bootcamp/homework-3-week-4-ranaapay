package storageType

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	AuthorId    uint   `json:"authorId"`
	Price       string `json:"price"`
	PublishTime string `json:"publishTime"`
	ISBN        string `json:"isbn"`
	PageNumber  int    `json:"pageNumber"`
	StockCode   string `json:"stockCode"`
	StockNumber int    `json:"stockNumber"`
}

func (b *Book) GetBookId() uint {
	return b.ID
}
func (b *Book) GetBookName() string {
	return b.Name
}
func (b *Book) GetAuthorId() uint {
	return b.AuthorId
}
func (b *Book) GetPublishTime() string {
	return b.PublishTime
}
func (b *Book) GetPrice() string {
	return b.Price
}
func (b *Book) GetPageNumber() int {
	return b.PageNumber
}
func (b *Book) GetISBN() string {
	return b.ISBN
}
func (b *Book) GetStockCode() string {
	return b.StockCode
}
func (b *Book) GetDeletedTime() gorm.DeletedAt {
	return b.DeletedAt
}
func (b *Book) GetStockNumber() int {
	return b.StockNumber
}
func (b *Book) SetStockNumber(num int) {
	b.StockNumber = num
}
func (b *Book) SetDeletedTime(isDeleted bool) {
	b.DeletedAt.Time = time.Now()
}

func (b *Book) PrintBook() {
	fmt.Printf("Book : %d ---------------\n"+
		"Name : %s\n"+
		"AuthorId : %d\n"+
		"Publish Year : %s\n"+
		"Price : %s\n"+
		"Page Number : %d\n"+
		"ISBN : %s\n"+
		"Stock Code : %s\n"+
		"StockNumber : %d\n"+
		"Deleted : %s\n",
		b.GetBookId(), b.GetBookName(), b.GetAuthorId(), b.GetPublishTime(), b.GetPrice(),
		b.GetPageNumber(), b.GetISBN(), b.GetStockCode(), b.GetStockNumber(), b.GetDeletedTime().Time.Format("2006-01-02 15:04:05"))
}
