package event

import (
	"encoding/json"
	"github.com/geiqin/supports/helper"
)

type MessageData struct {
	StoreId int64
	Data    interface{}
}

func (e MessageData) ToJson() string {
	return helper.JsonEncode(e)
}

func (e MessageData) Scan(jsonString string, dist interface{}) interface{} {
	err := json.Unmarshal([]byte(jsonString), dist)
	if err != nil {
		return nil
	}
	return dist
}
