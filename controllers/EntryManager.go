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
	// reading path variable
	a := "kevin"

	result := model.JsonToMap()
	entryarry := result["get"].([]interface{})
	value := make(map[string]interface{})

	for i, _ := range entryarry {
		mapentrys := entryarry[i].(map[string]interface{})
		for k, v := range mapentrys {
			if k == "name" && v == a {
				fmt.Println(" i have found the value")
				value = mapentrys
				break
			}
		}
	}
	fmt.Println("The value %v", value)
	jsonstring, _ := json.Marshal(value)

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonstring))

}
