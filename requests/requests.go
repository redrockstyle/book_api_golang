package requests

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

type Authors struct {
	Name string `sql:"type varchar(50);not null;"`
	Age  int16  `sql:"type smallint;not null;"`
	gorm.Model
}

type Books struct {
	Name string `sql:"type varchar(50);not null;"`
	Year int16  `sql:"type smallint;not null;"`
	gorm.Model
}

type AuthorBookID struct {
	BookID   int16 `sql:"type smallint;not null;"`
	AuthorID int16 `sql:"type smallint;not null;"`
}

func GetBooks(offset, limit string, Db *gorm.DB) any {

	if len(offset) == 0 || len(limit) == 0 {
		log.Printf("get book empty params")
		return []byte("")
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		log.Printf("%v: get book offset", err)
		return []byte("")
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Printf("%v: get book limit", err)
		return []byte("")
	}

	var result []Books
	res := Db.Offset(offsetInt).Limit(limitInt).Find(&result)
	if res.Error != nil {
		log.Printf("%v: get book json", res.Error)
		return []byte("")
	}
	return result
}

func GetBookOfId(id string, Db *gorm.DB) any {

	if len(id) == 0 {
		log.Printf("get book empty params")
		return []byte("")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("%v: get book id atoi", err)
		return []byte("")
	}

	var result []Books
	res := Db.Find(&result, idInt)
	if res.Error != nil {
		log.Printf("%v: get book id find", res.Error)
		return []byte("")
	}
	return result
}

func GetAuthors(offset, limit string, Db *gorm.DB) any {

	if len(offset) == 0 || len(limit) == 0 {
		log.Printf("get authors empty params")
		return []byte("")
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		log.Printf("%v: get authors offset", err)
		return []byte("")
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Printf("%v: get authors limit", err)
		return []byte("")
	}

	var result []Authors
	res := Db.Offset(offsetInt).Limit(limitInt).Find(&result)
	if res.Error != nil {
		log.Printf("%v: get authors json", res.Error)
		return []byte("")
	}
	return result
}

func GetAuthorOfId(id string, Db *gorm.DB) any {

	if len(id) == 0 {
		log.Printf("get author empty params")
		return []byte("")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("%v: get author id atoi", err)
		return []byte("")
	}

	var result []Authors
	res := Db.Find(&result, idInt)
	if res.Error != nil {
		log.Printf("%v: get author id find", res.Error)
		return []byte("")
	}
	return result
}

func AddBook(book, year string, Db *gorm.DB) any {

	if len(book) == 0 || len(year) == 0 {
		log.Printf("post book empty params")
		return []byte("")
	}
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		log.Printf("%v: post book atoi year", err)
		return []byte("")
	}

	bookStr := Books{Name: book, Year: int16(yearInt)}
	res := Db.Create(&bookStr)
	if res.Error != nil {
		log.Printf("%v: post book create", res.Error)
		return []byte("")
	}

	var resBook []Books
	res = Db.Find(&resBook, bookStr)
	if res.Error != nil {
		log.Printf("%v: post book json", res.Error)
		return []byte("")
	}
	return resBook
}

func AddAuthor(author, age string, Db *gorm.DB) any {

	if len(author) == 0 || len(age) == 0 {
		log.Printf("post author empty params")
		return []byte("")
	}
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		log.Printf("%v: post author atoi age", err)
		return []byte("")
	}

	authorStr := Authors{Name: author, Age: int16(ageInt)}
	res := Db.Create(&authorStr)
	if res.Error != nil {
		log.Printf("%v: post author create", res.Error)
		return []byte("")
	}

	var resAuthor []Authors
	res = Db.Find(&resAuthor, authorStr)
	if res.Error != nil {
		log.Printf("%v: error author json", res.Error)
		return []byte("")
	}
	return resAuthor
}

func BindID(idAuthor, idBook string, Db *gorm.DB) any {

	if len(idAuthor) == 0 || len(idBook) == 0 {
		log.Printf("post bind empty params")
		return []byte("")
	}
	intIdAuthor, err := strconv.Atoi(idAuthor)
	if err != nil {
		log.Printf("%v: post bind atoi author", err)
		return []byte("")
	}
	intIdBook, err := strconv.Atoi(idBook)
	if err != nil {
		log.Printf("%v: post bind atoi book", err)
		return []byte("")
	}

	idStr := AuthorBookID{AuthorID: int16(intIdAuthor), BookID: int16(intIdBook)}
	res := Db.Create(&idStr)
	if res.Error != nil {
		log.Printf("%v: post bind create", res.Error)
		return []byte("")
	}

	var resABid []AuthorBookID
	res = Db.Find(&resABid, idStr)
	if res.Error != nil {
		log.Printf("%v: post bind json", res.Error)
		return []byte("")
	}
	return resABid
}

func UpdateBook(id, book, year string, Db *gorm.DB) any {

	if len(id) == 0 || len(book) == 0 || len(year) == 0 {
		log.Printf("get book empty params")
		return []byte("")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("%v: update book atoi id", err)
		return []byte("")
	}
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		log.Printf("%v: update book atoi year", err)
		return []byte("")
	}

	bookStr := Books{Name: book, Year: int16(yearInt)}
	res := Db.Find(&Books{}, idInt).Update(bookStr)
	if res.Error != nil {
		log.Printf("%v: update book update", res.Error)
		return []byte("")
	}

	var resBook []Books
	res = Db.Find(&resBook, bookStr)
	if res.Error != nil {
		log.Printf("%v: update book json", res.Error)
		return []byte("")
	}
	return resBook
}

func UpdateAuthor(id, author, age string, Db *gorm.DB) any {

	if len(id) == 0 || len(author) == 0 || len(age) == 0 {
		log.Printf("get book empty params")
		return []byte("")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("%v: update book atoi id", err)
		return []byte("")
	}
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		log.Printf("%v: update book atoi year", err)
		return []byte("")
	}

	bookStr := Authors{Name: author, Age: int16(ageInt)}
	res := Db.Find(&Authors{}, idInt).Update(bookStr)
	if res.Error != nil {
		log.Printf("%v: update author find", res.Error)
		return []byte("")
	}

	var resAuthor []Authors
	res = Db.Find(&resAuthor, bookStr)
	if res.Error != nil {
		log.Printf("%v: update author json", res.Error)
		return []byte("")
	}
	return resAuthor
}

func DeleteBook(id string, Db *gorm.DB) any {

	if len(id) == 0 {
		log.Printf("delete book empty params")
		return []byte("")
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("%v: delete book atoi", err)
		return []byte("")
	}

	var resBook []Books
	res := Db.Find(&resBook, intId)
	if res.Error != nil {
		log.Printf("%v: delete book find", res.Error)
		return []byte("")
	}
	res = Db.Unscoped().Delete(&AuthorBookID{}, fmt.Sprintf("book_id = %v", intId))
	res = Db.Unscoped().Delete(&Books{}, intId)
	if res.Error != nil {
		log.Printf("%v: delete book delete", res.Error)
		return []byte("")
	}
	return resBook
}

func DeleteAuthor(id string, Db *gorm.DB) any {

	if len(id) == 0 {
		log.Printf("delete author empty params")
		return []byte("")
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("%v: delete author atoi", err)
		return []byte("")
	}

	var resAuthor []Authors
	res := Db.Find(&resAuthor, intId)
	if res.Error != nil {
		log.Printf("%v: delete author find", res.Error)
		return []byte("")
	}
	res = Db.Unscoped().Delete(&AuthorBookID{}, fmt.Sprintf("author_id = %v", intId))
	res = Db.Unscoped().Delete(&Authors{}, intId)
	if res.Error != nil {
		log.Printf("%v: delete author delete", res.Error)
		return []byte("")
	}
	return resAuthor
}
