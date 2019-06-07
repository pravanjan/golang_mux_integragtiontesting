package model

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"path"
	"os"
	"encoding/json"
)

func JsonToMap() map[string]interface{} {
	
	jsonFile, err := os.Open(getfilePath())
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

    var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	fmt.Printf("%v",result)
	return result
}
func getfilePath()string{
	_,filename,_,_ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "data.json")
	fmt.Println(filepath)
	return filepath
}
