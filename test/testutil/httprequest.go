package testutil

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
)

func CreateRequestContext(method string, url string, payload string,
	handler http.HandlerFunc) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	newhandler := http.HandlerFunc(handler)
	if payload != "" {
		reader := strings.NewReader(payload)
		req, _ := http.NewRequest(method, url, reader)
		newhandler.ServeHTTP(res, req)
	} else {
		req, _ := http.NewRequest(method, url, nil)
		newhandler.ServeHTTP(res, req)
	}

	return res

}

func CreateRequestContextPath(method string, url string, payload string,
	handler http.HandlerFunc, value []string) *httptest.ResponseRecorder {

	router := mux.NewRouter()
	router.HandleFunc(url, handler)

	res := httptest.NewRecorder()
	path := replacePathwithRealValue(url, value)
	fmt.Println(path)
	if payload != "" {
		req, _ := http.NewRequest(method, path, nil)
		router.ServeHTTP(res, req)

	} else {
		req, _ := http.NewRequest(method, path, nil)
		router.ServeHTTP(res, req)

	}

	return res

}

func replacePathwithRealValue(url string, value []string) string {
	fmt.Println(value)

	for index, element := range value {
		fmt.Println(index, element)
		strEx := url
		reStr := regexp.MustCompile("^(.*?)({.*?})(.*)$")
		repStr := "${1}" + element + "${3}"
		url = reStr.ReplaceAllString(strEx, repStr)
	}
	fmt.Println(url)
	return url
}
