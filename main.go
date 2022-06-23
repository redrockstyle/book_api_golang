package main

import (
	"book_api/database"
	"book_api/requests"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type BookForm struct {
	Book string `json:"book"`
	Year string `json:"year"`
}

type AuthorForm struct {
	Author string `json:"author"`
	Age    string `json:"age"`
}

type ABIDForm struct {
	AID string `json:"aid"`
	BID string `json:"bid"`
}

type UpdateBookForm struct {
	Id   string `json:"id"`
	Book string `json:"book"`
	Year string `json:"year"`
}

type UpdateAuthorForm struct {
	Id     string `json:"id"`
	Author string `json:"author"`
	Age    string `json:"age"`
}

type OffsetLimitForm struct {
	Offset string `json:"offset"`
	Limit  string `json:"limit"`
}

func main() {
	Db := database.InitDatabase()
	database.InitTable(Db)
	database.SomethingInputDB(Db)
	InitREST(Db)
}

func InitREST(Db *gorm.DB) {
	app := fiber.New()
	// GET
	app.Get("/book/all", func(c *fiber.Ctx) error {
		olForm := new(OffsetLimitForm)
		if err := c.BodyParser(olForm); err != nil {
			log.Printf("%v: get all book", err)
			return err
		}
		books := requests.GetBooks(olForm.Offset, olForm.Limit, Db)
		return c.JSON(books)
	})
	app.Get("/author/all", func(c *fiber.Ctx) error {
		olForm := new(OffsetLimitForm)
		if err := c.BodyParser(olForm); err != nil {
			log.Printf("%v: get all authors", err)
			return err
		}
		authors := requests.GetAuthors(olForm.Offset, olForm.Limit, Db)
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
		authorForm := new(AuthorForm)
		if err := c.BodyParser(authorForm); err != nil {
			return err
		}
		addAuthor := requests.AddAuthor(authorForm.Author, authorForm.Age, Db)
		return c.JSON(addAuthor)
	})
	app.Post("/book", func(c *fiber.Ctx) error {
		bookForm := new(BookForm)
		if err := c.BodyParser(bookForm); err != nil {
			return err
		}
		addBook := requests.AddBook(bookForm.Book, bookForm.Year, Db)
		return c.JSON(addBook)
	})
	app.Post("/bind", func(c *fiber.Ctx) error {
		bookForm := new(ABIDForm)
		if err := c.BodyParser(bookForm); err != nil {
			return err
		}
		addBook := requests.BindID(bookForm.AID, bookForm.BID, Db)
		return c.JSON(addBook)
	})

	// PUT
	app.Put("/book/update", func(c *fiber.Ctx) error {
		upbForm := new(UpdateBookForm)
		if err := c.BodyParser(upbForm); err != nil {
			return err
		}
		updateBook := requests.UpdateBook(upbForm.Id, upbForm.Book, upbForm.Year, Db)
		return c.JSON(updateBook)
	})
	app.Put("/author/update", func(c *fiber.Ctx) error {
		upaForm := new(UpdateAuthorForm)
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
