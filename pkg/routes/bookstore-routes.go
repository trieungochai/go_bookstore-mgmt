package routes

import (
	"github.com/gorilla/mux"
	"github.com/trieungochai/go_bookstore-mgmt/pkg/controllers"
)

var RegisterBookstoreRoutes = func(router *mux.Route) {
	router.HandlerFunc("/books/", controllers.CreateBook).Methods("POST")
	router.HandlerFunc("/books/", controllers.GetAllBooks).Methods("GET")
	router.HandlerFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandlerFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandlerFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
