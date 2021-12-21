package utils

import (
	"errors"
	"fmt"
)

type JSONObject struct {
	Value string `json:"value"`
}

func (jsono *JSONObject) UnmarshalJSON(data []byte) error {
	fmt.Println(len(data))
	if data[0] != 123 || data[len(data)-1] != 125 {
		message := fmt.Sprintf("invalid json object: %s", string(data))
		return errors.New(message)
	}
	jsono.Value = string(data)
	return nil
}
