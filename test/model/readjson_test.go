package model
import (
	"testing"
	"integration/testing/model"
	"fmt"
)


func TestJsonToMap_NotEmpty(t *testing.T) {
	result := model.JsonToMap()

	if result!=nil{
		fmt.Println(result)
	}
}
func TestJsonToMap_validateAllResultType(t *testing.T) {
	result := model.JsonToMap()
	
	createMap := result["create"].(map[string]interface{})
	
	if createMap["name"]=="kunal" {
		fmt.Println(createMap["name"])
	}
	
	entryList := result["get"].([]interface{})
	
	if entryList!=nil {
		fmt.Println(entryList[0])

	}



}