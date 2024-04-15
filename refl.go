package goutilities

import (
	"reflect"
)

func PriVal(val interface{}) (pri int64, err error) {
	if val == nil {
		return
	}

	switch rt := reflect.TypeOf(val).Kind(); rt {
	case reflect.Int:
		pri = int64(val.(int))
	case reflect.Int8:
		pri = int64(val.(int8))
	case reflect.Int32:
		pri = int64(val.(int32))
	case reflect.Int64:
		pri = val.(int64)
	default:
		return pri, err
	}

	return pri, nil
}

// Converts any key value to int64. Accepts interface{}
// and type-cast to int64
func AnyToInt64(val interface{}) (pri int64, err error) {
	if val == nil {
		return
	}

	switch rt := reflect.TypeOf(val).Kind(); rt {
	case reflect.Int:
		pri = int64(val.(int))
	case reflect.Int8:
		pri = int64(val.(int8))
	case reflect.Int32:
		pri = int64(val.(int32))
	case reflect.Int64:
		pri = val.(int64)
  case reflect.Float32:
    pri = int64(val.(float32))
  case reflect.Float64:
    pri = int64(val.(float64))
	default:
		return pri, err
	}

	return pri, nil
}
