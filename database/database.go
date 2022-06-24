package database

import (
	"book_api/requests"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"net/url"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Path     string
}

func SomethingInputDB(Db *gorm.DB) {
	Db.Create(&requests.Authors{Name: "Магомед", Age: 18})
	Db.Create(&requests.Authors{Name: "Ахмед", Age: 19})
	Db.Create(&requests.Authors{Name: "Абдул", Age: 22})

	Db.Create(&requests.Books{Name: "Братская иситина Магомеда", Year: 2015})
	Db.Create(&requests.Books{Name: "Братская иситина Ахмеда", Year: 2017})
	Db.Create(&requests.Books{Name: "Братская иситина Абдула", Year: 2018})

	Db.Create(&requests.AuthorBookID{AuthorID: 1, BookID: 1})
	Db.Create(&requests.AuthorBookID{AuthorID: 2, BookID: 2})
	Db.Create(&requests.AuthorBookID{AuthorID: 3, BookID: 3})

	Db.Create(&requests.Users{Login: "name", Pass: []byte("pass")})
}

func InitDatabase(cfg DBConfig) *gorm.DB {
	var err error
	dsn := url.URL{
		User:     url.UserPassword(cfg.User, cfg.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%v:%v", cfg.Host, cfg.Port),
		Path:     cfg.Path,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	Db, err := gorm.Open("postgres", dsn.String())
	if err != nil {
		log.Fatal(err)
	}
	return Db
}

func InitTable(Db *gorm.DB) {
	Db.DropTable(&requests.Authors{})
	Db.DropTable(&requests.Books{})
	Db.DropTable(&requests.AuthorBookID{})
	Db.DropTable(&requests.Users{})

	Db.CreateTable(&requests.Authors{})
	Db.CreateTable(&requests.Books{})
	Db.CreateTable(&requests.AuthorBookID{})
	Db.CreateTable(&requests.Users{})
}
