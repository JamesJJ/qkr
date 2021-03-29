package main

import (
	"encoding/base64"
	"encoding/json"
)

func JSON(v interface{}) string {
	return string(JSONBytes(v))
}

func JSONBase64(v interface{}) string {
	return base64.StdEncoding.EncodeToString(JSONBytes(v))
}

func JSONBytes(v interface{}) []byte {
	if j, e := json.Marshal(v); e == nil {
		return j
	}
	return []byte("{}")
}
