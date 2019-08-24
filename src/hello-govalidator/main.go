package main

import (
	"fmt"
	"hello-govalidator/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	
	// Route handles & endpoints
	router.HandleFunc("/tokens", controllers.GetTokens).Methods("GET", "POST")
	

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
