package controllers

import (
	"encoding/json"
	"fmt"
	"integration/testing/controllers"
	"integration/testing/model"
	"integration/testing/test/testutil"
	"net/http"
	"testing"
)

func TestGetAllEntry(t *testing.T) {
	res := testutil.CreateRequestContext("GET", "/api/entry/get", "", controllers.GetAllEntry)

	if status := res.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}
	// convert the resonse body to map to read the value
	var arrayMap []map[string]interface{}
	json.Unmarshal([]byte(res.Body.String()), &arrayMap)
	if arrayMap[0]["name"] != nil {
		fmt.Println(arrayMap[0]["name"])
	}

}

func TestCreateNewEntry(t *testing.T) {
	//getting post data from json file
	data := model.JsonToMap()
	payload := data["create"].(map[string]interface{})
	jsonstring, _ := json.Marshal(payload)
	res := testutil.CreateRequestContext("GET", "/api/entry/create", string(jsonstring), controllers.CreateNewEntry)

	if status := res.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}
	// convert the resonse body to map to read the value
	if res.Body.String() != "created successfully" {
		t.Errorf("unexpected response")
	}
}

func TestSearchEntry(t *testing.T) {
	name := []string{"kevin"}
	res := testutil.CreateRequestContextPath("GET", "/api/entry/{name}/search/", "", controllers.SearchEntry, name)
	fmt.Println("*****" + res.Body.String())
	if status := res.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}
	//un marshall the data

	value := make(map[string]string)
	json.Unmarshal([]byte(res.Body.String()), &value)
	if value["role"] != "developer" {
		t.Errorf("do not receive correct value")
	}

}
