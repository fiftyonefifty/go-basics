package main

import (
	"fmt"
	"hello-rest/controllers"
	"hello-rest/models"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	//router.NotFoundHandler = app.NotFoundHandler
	// Hardcoded data - @todo: add database
	controllers.Books = append(controllers.Books, models.Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &models.Author{Firstname: "John", Lastname: "Doe"}})
	controllers.Books = append(controllers.Books, models.Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &models.Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route handles & endpoints
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}
}
