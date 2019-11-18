package util

import (
	"encoding/json"
	"reflect"
)

func JsonToObject(data []byte, v interface{}) {
	json.Unmarshal(data, &v)
}

func ObjectToJson(v interface{}) []byte {
	if v == nil || reflect.ValueOf(v).IsNil() {
		return []byte("{}")
	}
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return []byte("{\"error\":\"Json Fail\"}")
	}
	return data
}
