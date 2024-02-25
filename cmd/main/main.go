package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"bookstore/pkg/routes"
)

func main(){
	res := mux.NewRouter()
	routes.RegisterBookStoreRoutes(res)
	http.Handle("/", res)
	log.Fatal(http.ListenAndServe("localhost:3306", res))
}