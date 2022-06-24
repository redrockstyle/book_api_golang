package main

import (
	"book_api/database"
	"book_api/requests"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func main() {
	cfg := database.DBConfig{User: "user", Password: "password", Host: "192.168.224.130", Port: "5432", Path: "books"}

	Db := database.InitDatabase(cfg)
	database.InitTable(Db)
	database.SomethingInputDB(Db)
	InitREST(Db)
}

func InitREST(Db *gorm.DB) {
	app := fiber.New()

	// LOGIN
	app.Post("/login", func(c *fiber.Ctx) error {
		userForm := new(requests.UserForm)
		if err := c.BodyParser(userForm); err != nil {
			log.Printf("%v: post login", err)
			return err
		}
		user := requests.Login(userForm.Login, userForm.Password, Db)
		return c.JSON(user)
	})
	app.Post("/registration", func(c *fiber.Ctx) error {
		userForm := new(requests.UserForm)
		if err := c.BodyParser(userForm); err != nil {
			log.Printf("%v: post registrtion", err)
			return err
		}
		user := requests.Registration(userForm.Login, userForm.Password, Db)
		return c.JSON(user)
	})

	// GET
	app.Get("/book/all", func(c *fiber.Ctx) error {
		books := requests.GetBooks(c.Query("offset"), c.Query("limit"), Db)
		return c.JSON(books)
	})
	app.Get("/author/all", func(c *fiber.Ctx) error {
		authors := requests.GetAuthors(c.Query("offset"), c.Query("limit"), Db)
		return c.JSON(authors)
	})
	app.Get("/book/:id", func(c *fiber.Ctx) error {
		book := requests.GetBookOfId(c.Params("id"), Db)
		return c.JSON(book)
	})
	app.Get("/author/:id", func(c *fiber.Ctx) error {
		author := requests.GetAuthorOfId(c.Params("id"), Db)
		return c.JSON(author)
	})

	// POST
	app.Post("/author", func(c *fiber.Ctx) error {
		authorForm := new(requests.AuthorForm)
		if err := c.BodyParser(authorForm); err != nil {
			return err
		}
		addAuthor := requests.AddAuthor(authorForm.Author, authorForm.Age, Db)
		return c.JSON(addAuthor)
	})
	app.Post("/book", func(c *fiber.Ctx) error {
		bookForm := new(requests.BookForm)
		if err := c.BodyParser(bookForm); err != nil {
			return err
		}
		addBook := requests.AddBook(bookForm.Book, bookForm.Year, Db)
		return c.JSON(addBook)
	})
	app.Post("/bind", func(c *fiber.Ctx) error {
		bookForm := new(requests.ABIDForm)
		if err := c.BodyParser(bookForm); err != nil {
			return err
		}
		addBook := requests.BindID(bookForm.AID, bookForm.BID, Db)
		return c.JSON(addBook)
	})

	// PUT
	app.Put("/book/update", func(c *fiber.Ctx) error {
		upbForm := new(requests.UpdateBookForm)
		if err := c.BodyParser(upbForm); err != nil {
			return err
		}
		updateBook := requests.UpdateBook(upbForm.Id, upbForm.Book, upbForm.Year, Db)
		return c.JSON(updateBook)
	})
	app.Put("/author/update", func(c *fiber.Ctx) error {
		upaForm := new(requests.UpdateAuthorForm)
		if err := c.BodyParser(upaForm); err != nil {
			return err
		}
		updateAuthor := requests.UpdateAuthor(upaForm.Id, upaForm.Author, upaForm.Age, Db)
		return c.JSON(updateAuthor)
	})

	// DELETE
	app.Delete("/book/:id", func(c *fiber.Ctx) error {
		deleteBook := requests.DeleteBook(c.Params("id"), Db)
		return c.JSON(deleteBook)
	})
	app.Delete("/author/:id", func(c *fiber.Ctx) error {
		deleteAuthor := requests.DeleteAuthor(c.Params("id"), Db)
		return c.JSON(deleteAuthor)
	})

	// SIN SOBAKI
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("сын собаки")
	})

	log.Fatal(app.Listen(":3000"))
}
