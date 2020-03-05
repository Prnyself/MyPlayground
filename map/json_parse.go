package _map

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Unmarshal(input map[string]interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	fmt.Println(m)
	fmt.Println(reflect.TypeOf(m["type"]))
	return nil
}
