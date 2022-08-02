package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alireza/bookstore/internal/models"
	"github.com/alireza/bookstore/internal/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetAllBooks(writer http.ResponseWriter, request *http.Request) {
	allBooks := models.GetAllBooks()
	res, _ := json.Marshal(allBooks)
	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}
func GetBookByid(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookID := vars["id"]
	id, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book, _ := models.GetBookByid(id)
	res, _ := json.Marshal(book)
	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}
func CreateBook(writer http.ResponseWriter, request *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(request, newBook)
	b := newBook.CreateBook()
	res, _ := json.Marshal(b)
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}
func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookID := vars["id"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}
func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	bookUpdate := &models.Book{}
	utils.ParseBody(request, bookUpdate)
	vars := mux.Vars(request)
	bookID := vars["id"]
	ID, err := strconv.ParseInt(bookID, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, database := models.GetBookByid(ID)
	if bookUpdate.Name != "" {
		bookDetails.Name = bookUpdate.Name
	}
	if bookUpdate.Author != "" {
		bookDetails.Author = bookUpdate.Author
	}
	if bookUpdate.Publication != "" {
		bookDetails.Publication = bookUpdate.Publication
	}
	database.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}
