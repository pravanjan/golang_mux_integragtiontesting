package controllers

import (
	"encoding/json"
	"fmt"
	"integration/testing/model"
	"net/http"
)

var CreateNewEntry = func(writer http.ResponseWriter, request *http.Request) {
	jsonvalue := json.NewDecoder(request.Body)
	fmt.Printf("%v", jsonvalue)
	fmt.Fprint(writer, "created successfully")

}
var GetAllEntry = func(w http.ResponseWriter, r *http.Request) {
	result := model.JsonToMap()
	entryarry := result["get"].([]interface{})
	fmt.Printf("%v", entryarry)

	jsonstring, _ := json.Marshal(entryarry)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonstring))
}

var SearchEntry = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("default view ")

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "search successfully")

}
