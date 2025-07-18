package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type ChartSetting struct {
	X int
	Y int
}

type FormSetting struct {
	FormId string
}

func simple(jsonString string) {
	var ip map[string]any
	json.Unmarshal([]byte(jsonString), &ip)

	switch ip["type"].(string) {
	case "chart":
		{
			var res ChartSetting
			json.Unmarshal([]byte(jsonString), &res)
			fmt.Println(res)
		}
	case "form":
		{
			var res FormSetting
			json.Unmarshal([]byte(jsonString), &res)
			fmt.Println(res)
		}
	}
}

func notSimple(jsonString string) {
	var ip map[string]any
	json.Unmarshal([]byte(jsonString), &ip)

	switch ip["type"].(string) {
	case "chart":
		{
			var res ChartSetting
			mapstructure.Decode(ip, &res)
			fmt.Println(res)
		}
	case "form":
		{
			var res FormSetting
			mapstructure.Decode(ip, &res)
			fmt.Println(res)
		}
	}
}

func main() {
	jsonString := `{
    "type": "chart",
    "setting": {
      "x": 25,
      "y": 25
    }
  }`

	simple(jsonString)
	notSimple(jsonString)
}
