package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type ResultJSON struct {
	Status     int      `json:"status"`
	Disease    []string `json:"disease"`
	ResponseID string   `json:"response_id"`
}

func structToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(data)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := typ.Field(i).Tag.Get("json")
		result[tag] = field.Interface()
	}

	return result
}

func main() {
	// 模拟一个 resultJSON
	jsonData := []byte(`{
		"status": 0,
		"disease": ["普通感冒", "急性胃肠炎"],
		"response_id": "22994434614890496"
	}`)

	var resultJSON ResultJSON
	err := json.Unmarshal(jsonData, &resultJSON)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 转换为字典
	// resultMap := map[string]interface{}{
	// 	"status":      resultJSON.Status,
	// 	"disease":     resultJSON.Disease,
	// 	"response_id": resultJSON.ResponseID,
	// }
	resultMap := structToMap(resultJSON)
	// 打印resultJSON
	fmt.Printf("%#v\n", resultJSON)
	// 打印字典
	fmt.Printf("%#v\n", resultMap)
}
