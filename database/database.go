package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"net/url"
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

func SomethingInputDB(Db *gorm.DB) {
	user := Authors{Name: "Магомед", Age: 18}
	user2 := Authors{Name: "Ахмед", Age: 18}
	user3 := Authors{Name: "Абдул", Age: 18}

	result := Db.Create(&user)
	fmt.Printf("AUTHORS -> ID:%v Error:%v RowsAffected:%v\n", user.ID, result.Error, result.RowsAffected)
	result2 := Db.Create(&user2)
	fmt.Printf("AUTHORS -> ID:%v Error:%v RowsAffected:%v\n", user2.ID, result2.Error, result2.RowsAffected)
	result3 := Db.Create(&user3)
	fmt.Printf("AUTHORS -> ID:%v Error:%v RowsAffected:%v\n", user3.ID, result3.Error, result3.RowsAffected)

	book := Books{Name: "Братская иситина Магомеда", Year: 2010}
	book1 := Books{Name: "Братская иситина Ахмеда", Year: 2011}
	book2 := Books{Name: "Братская иситина Абдула", Year: 2012}

	result4 := Db.Create(&book)
	fmt.Printf("AUTHORS -> ID:%v Error:%v RowsAffected:%v\n", book.ID, result4.Error, result4.RowsAffected)
	result5 := Db.Create(&book1)
	fmt.Printf("AUTHORS -> ID:%v Error:%v RowsAffected:%v\n", book1.ID, result5.Error, result5.RowsAffected)
	result6 := Db.Create(&book2)
	fmt.Printf("AUTHORS -> ID:%v Error:%v RowsAffected:%v\n", book2.ID, result6.Error, result6.RowsAffected)

	authBookId := AuthorBookID{AuthorID: 1, BookID: 1}
	authBookId1 := AuthorBookID{AuthorID: 2, BookID: 2}
	authBookId2 := AuthorBookID{AuthorID: 3, BookID: 3}

	result7 := Db.Create(&authBookId)
	fmt.Printf("AUTHORS -> ID:%v Error:%v RowsAffected:%v\n", authBookId.BookID, result7.Error, result7.RowsAffected)
	result8 := Db.Create(&authBookId1)
	fmt.Printf("AUTHORS -> ID:%v Error:%v RowsAffected:%v\n", authBookId1.BookID, result8.Error, result8.RowsAffected)
	result9 := Db.Create(&authBookId2)
	fmt.Printf("AUTHORS -> ID:%v Error:%v RowsAffected:%v\n", authBookId2.BookID, result9.Error, result9.RowsAffected)
}

func InitDatabase() *gorm.DB {
	var err error
	dsn := url.URL{
		User:     url.UserPassword("user", "password"),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%v:%v", "192.168.224.130", "5432"),
		Path:     "books",
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	//fmt.Println(dsn.String())
	Db, err := gorm.Open("postgres", dsn.String())
	if err != nil {
		log.Fatal(err)
	}
	return Db
}

func InitTable(Db *gorm.DB) {
	Db.DropTable(&Authors{})
	Db.DropTable(&Books{})
	Db.DropTable(&AuthorBookID{})

	Db.CreateTable(&Authors{})
	Db.CreateTable(&Books{})
	Db.CreateTable(&AuthorBookID{})
}
