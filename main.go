package main

import (
	"book-store/controllers"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var ctx context.Context
var err error
var client *mongo.Client
var MongoUri string = "mongodb://tanveeshs:pass123@localhost:27017/book-server?authSource=admin"
var bookController *controllers.BookController

func init() {
	ctx = context.Background()
	client, err = mongo.Connect(ctx,
		options.Client().ApplyURI(MongoUri))
	if err = client.Ping(context.TODO(),
		readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	collection := client.Database("book-server").Collection("books")
	bookController = controllers.NewBookController(collection,
		ctx)
}
func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/books", bookController.GetBooks)
	app.Get("/books/all", bookController.GetAllBooks)
	app.Post("/books", bookController.AddBook)
	app.Get("/books/:id", bookController.GetBookById)
	app.Put("/books/:id", bookController.UpdateBook)
	app.Delete("/books/:id", bookController.DeleteBook)
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Error in running the server")
		return
	}
	log.Println("Server is running")
}
