### Introduction:
The Book Store Application is a web-based service that provides a set of APIs for managing and retrieving information about books. It uses the Fiber web framework for handling HTTP requests and communicates with a MongoDB database for data storage.

### Dependencies:
The application relies on the following external libraries and packages:
- Fiber (github.com/gofiber/fiber/v2): A fast and lightweight web framework for handling HTTP requests.
- MongoDB Go Driver (go.mongodb.org/mongo-driver): A driver for connecting to a MongoDB database.


### MongoDB Configuration:
The application connects to a MongoDB database using the provided connection string. The database is named "book-server," and it contains a collection named "books." The MongoDB connection string and database configuration are specified in the main.go file.

### API Endpoints:
The following API endpoints are defined in the application:

- GET /books: Retrieve a list of published books.
- GET /books/all: Retrieve a list of all books, including unpublished ones.
- POST /books: Add a new book to the database.
- GET /books/:id: Retrieve information about a specific book by its unique ID.
- PUT /books/:id: Update information about a specific book.
- DELETE /books/:id: Delete a book by its unique ID.

### Data Models:
The application works with a data model named Book, which is defined in the models package. A Book has the following properties:

- ID (ObjectID): A unique identifier for the book.
- Name (string): The name of the book.
- Author (string): The author of the book.
- Genre (string): The genre of the book.
- ISBN (string): The International Standard Book Number.
- Published (bool): A flag indicating whether the book is published.

### Request and Response Formats:
- For adding a book (POST /books), the request body should contain a JSON object with the name, author, genre, and isbn fields.
- For updating a book (PUT /books/:id), the request body should contain a JSON object with fields to be updated, including name, author, genre, isbn, and published.

### Error Handling:
The application handles errors gracefully and returns appropriate HTTP status codes and error messages when necessary. For example, it returns a 400 Bad Request status code for invalid requests and a 404 Not Found status code when a requested book does not exist.

### Trade-offs for Choosing MongoDB:
MongoDB was chosen as the database for this application for the following reasons:

- Flexible Schema: MongoDB is a NoSQL database that supports flexible schemas. This flexibility is beneficial when dealing with data like books, which can have varying attributes.

- Scalability: MongoDB can handle large amounts of data and allows for easy horizontal scaling, making it suitable for applications that may experience rapid growth.

- JSON-like Documents: MongoDB stores data in JSON-like documents, which aligns well with the data format typically used in web applications (e.g., JSON).


### API Documentation

https://documenter.getpostman.com/view/8565014/2s9YRCYBj6