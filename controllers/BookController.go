package controllers

import (
	"book-store/models"
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type BookController struct {
	Collection *mongo.Collection
	Ctx        context.Context
}

func NewBookController(collection *mongo.Collection, ctx context.Context) *BookController {
	return &BookController{
		Collection: collection,
		Ctx:        ctx,
	}
}

type AddBookReq struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Isbn   string `json:"isbn"`
}
type UpdateBookReq struct {
	Name      string `json:"name"`
	Author    string `json:"author"`
	Genre     string `json:"genre"`
	Isbn      string `json:"isbn"`
	Published bool   `json:"published"`
}

func (handler *BookController) GetBooks(c *fiber.Ctx) error {
	query := bson.M{"published": true}
	findOptions := options.Find().SetSort(bson.M{"created_at": 1})
	cur, err := handler.Collection.Find(handler.Ctx, query, findOptions)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}
	}(cur, handler.Ctx)
	books := make([]models.Book, 0)
	for cur.Next(handler.Ctx) {
		var book models.Book
		err := cur.Decode(&book)
		if err != nil {
			return err
		}
		books = append(books, book)
	}
	if len(books) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Not Found")
	} else {
		return c.JSON(fiber.Map{
			"books": books,
		})
	}
}

func (handler *BookController) GetAllBooks(c *fiber.Ctx) error {
	findOptions := options.Find().SetSort(bson.M{"created_at": 1})
	cur, err := handler.Collection.Find(handler.Ctx, bson.M{}, findOptions)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}
	}(cur, handler.Ctx)
	books := make([]models.Book, 0)
	for cur.Next(handler.Ctx) {
		var book models.Book
		err := cur.Decode(&book)
		if err != nil {
			return err
		}
		books = append(books, book)
	}
	if len(books) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Not Found")
	} else {
		return c.JSON(fiber.Map{
			"books": books,
		})
	}
}

func (handler *BookController) GetBookById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	log.Println("ID PARAM:", idParam)
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID format")
	}
	book := new(models.Book)
	err = handler.Collection.FindOne(handler.Ctx, bson.M{"_id": objectID, "published": true}).Decode(&book)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fiber.NewError(fiber.StatusNotFound, "Book not found")
		} else {
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}
	}
	return c.JSON(fiber.Map{
		"book": book,
	})
}

func (handler *BookController) DeleteBook(c *fiber.Ctx) error {
	idParam := c.Params("id")
	log.Println("ID PARAM:", idParam)
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID format")
	}
	_, err = handler.Collection.DeleteOne(handler.Ctx, bson.M{"_id": objectID})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fiber.NewError(fiber.StatusNotFound, "Book not found")
		} else {
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}
	}

	return c.JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}

func (handler *BookController) AddBook(c *fiber.Ctx) error {
	addBookReq := new(AddBookReq)
	if err := c.BodyParser(addBookReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	book := new(models.Book)
	book.ID = primitive.NewObjectID()
	book.CreatedAt = time.Now()
	book.Name = addBookReq.Name
	book.Author = addBookReq.Author
	book.Genre = addBookReq.Genre
	book.Isbn = addBookReq.Isbn
	book.Published = true
	savedBook, err := handler.Collection.InsertOne(handler.Ctx, book)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Unable to save user")
	}
	return c.JSON(fiber.Map{"message": "Book added", "id": savedBook.InsertedID})
}

func (handler *BookController) UpdateBook(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID format")
	}
	updateReq := new(UpdateBookReq)
	if err := c.BodyParser(updateReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	update := bson.M{}
	update["name"] = updateReq.Name
	update["author"] = updateReq.Author
	update["genre"] = updateReq.Genre
	update["isbn"] = updateReq.Isbn
	update["published"] = updateReq.Published
	result, err := handler.Collection.UpdateOne(handler.Ctx, bson.M{"_id": objectID}, bson.M{"$set": update})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update book")
	}
	if result.MatchedCount == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Book not found")
	}

	return c.JSON(fiber.Map{
		"message": "Book updated successfully",
	})
}
