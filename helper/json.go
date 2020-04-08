package helper

import "encoding/json"

func JsonDecode(jsonStr string, out interface{}) interface{} {
	json.Unmarshal([]byte(jsonStr), &out)
	return &out
}

func JsonEncode(obj interface{}) string {
	str, _ := json.Marshal(obj)
	return string(str)
}
