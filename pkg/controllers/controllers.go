package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Mahaveer86619/GO-Bookstore/pkg/models"
	"github.com/Mahaveer86619/GO-Bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {

	newbooks := models.GetBooks()

	res, _ := json.Marshal(newbooks)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)

	if err != nil {
		fmt.Println("Error while parsing. ", err)
	}

	bookDetails, _ := models.GetBookById(ID)

	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBooks(w http.ResponseWriter, r *http.Request) {

	book := &models.Book{}

	utils.ParseBody(r, book)

	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)

	if err != nil {
		fmt.Println("Error while parsing. ", err)
	}

	bookDetails := models.DeleteBook(ID)

	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	var book = &models.Book{}
	utils.ParseBody(r, book)

	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId,0,0)

	if err != nil {
		fmt.Println("Error while parsing. ", err)
	}

	bookDetails, db := models.GetBookById(ID)

	if book.Name != "" {
		bookDetails.Name = book.Name
	}
	if book.Author != "" {
		bookDetails.Author = book.Author
	}
	if book.Publication != "" {
		bookDetails.Publication = book.Publication
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	
}