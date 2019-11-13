package reflect

import (
	"errors"
	"fmt"
	"reflect"
)

func ShowTypeOf(input interface{}) error {
	fmt.Println(reflect.TypeOf(input).Kind())
	return nil
}

func SliceToInterface(input interface{}) ([]string, error) {
	if reflect.TypeOf(input).Kind() != reflect.Slice {
		return nil, errors.New("input invalid")
	}
	i, ok := input.([]string)
	if !ok {
		return nil, errors.New("input invalid")
	}

	return i, nil
}
