package controllers

import (
"fmt"
"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
"net/http"
"strconv"
"encoding/json"
"bookstore/pkg/utils"
"bookstore/pkg/models")

var NewBook models.Book

func CreateBook(sen http.ResponseWriter, res *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(res, CreateBook)
	b := CreateBook.CreateBook()
	r, _ := json.Marshal(b)
	sen.Header().Set("Content-Type", "pkglication/json")
	sen.WriteHeader(http.StatusOK)
	sen.Write(r)
}

func GetBook(sen http.ResponseWriter, res *http.Request) {
	newBook := models.GetAllBooks()
	r, _ := json.Marshal(newBook)
	// sen.Header().Set("Content-Type", "pkg/json")
	sen.Header().Set("Content-Type", "pkglication/json")
	sen.WriteHeader(http.StatusOK)
	sen.Write(r)
}

func GetBookById(sen http.ResponseWriter, res *http.Request) {
		vars  := mux.Vars(res)
		bookId := vars["bookId"]
		ID, err := strconv.ParseInt(bookId,0,0)
		if err != nil{
			fmt.Println("error while parsing")
		}
		bookDetails, _ := models.GetBookById(ID)
		r, _  := json.Marshal(bookDetails)
		sen.Header().Set("Content-Type", "pkglication/json")
		sen.WriteHeader(http.StatusOK)
		sen.Write(r)
}

func DeleteBook(sen http.ResponseWriter, res *http.Request) {
del := mux.Vars(res)
deleteId := del["deleteId"]
ID, err := strconv.ParseInt(deleteId,0,0)
if err != nil {
		fmt.Println("error while parsind delete")
	}
deleteBookDetails := models.DeleteBook(ID)
r,_ := json.Marshal(deleteBookDetails)
sen.Header().Set("Content-Type", "pkglication/json")
sen.WriteHeader(http.StatusOK)
sen.Write(r)
} 

func UpdateBook(sen http.ResponseWriter, res *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(res, updateBook)
	vars := mux.Vars(res)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err !=  nil{
		fmt.Println("error while passing")
	}
booksDetails, db := models.GetBookById(ID)

	if updateBook.Name != " "{
		booksDetails.Name = updateBook.Name
	}
	if updateBook.Author != " "{
		booksDetails.Author = updateBook.Author
	}
	if updateBook.Publication != " "{
		booksDetails.Publication = updateBook.Publication
	}
	db.Save(&booksDetails)
	r, _ := json.Marshal(booksDetails)
	sen.Header().Set("Content-Type", "pkglication/json")
	sen.WriteHeader(http.StatusOK)
	sen.Write(r)
}
