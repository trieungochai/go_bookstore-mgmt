package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/trieungochai/go_bookstore-mgmt/pkg/models"
	"github.com/trieungochai/go_bookstore-mgmt/pkg/utils"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	err := utils.ParseBody(r, newBook)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	book, err := newBook.CreateBook()
	if err != nil {
		http.Error(w, "failed to create book", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "error while  encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, "failed to fetch books", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "error while encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "invalid bookId", http.StatusBadRequest)
	}

	book, err := models.GetBookById(id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "error while encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "invalid bookId", http.StatusBadRequest)
		return
	}

	deletedBook, err := models.DeleteBook(id)
	if err != nil {
		http.Error(w, "Book not found or could not be deleted", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(deletedBook)
	if err != nil {
		http.Error(w, "error while encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatedBook := &models.Book{}
	err := utils.ParseBody(r, updatedBook)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "invalid bookId", http.StatusBadRequest)
		return
	}

	bookDetails, err := models.GetBookById(id)
	if err != nil {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}

	if updatedBook.Title != "" {
		bookDetails.Title = updatedBook.Title
	}
	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}

	if err := db.Save(&bookDetails).Error; err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	// Marshal the updated book details into JSON
	res, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
