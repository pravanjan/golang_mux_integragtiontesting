package main

import (
	"fmt"
	"log"
	"net/http"

	"integration/testing/controllers"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("inside ds method ")
	router := mux.NewRouter()
	router.HandleFunc("/api/entry/get", controllers.GetAllEntry)
	router.HandleFunc("/api/entry/create/", controllers.CreateNewEntry).Methods("POST")
	router.HandleFunc("/api/entry/{name}/search/", controllers.SearchEntry).Methods("GET")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8888", router))

}
