package tools

import "encoding/json"

func ToJson(v interface{}) string {
	str, _ := json.Marshal(v)
	return string(str)
}
