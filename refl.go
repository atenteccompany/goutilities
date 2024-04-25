package goutilities

import (
	"fmt"
	"reflect"
	"strings"
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

// ---------------------------------------------------------------------------------------------------------
type fieldInfo struct {
	Name  string
	Type  string
	Index int
}

func PrintWithReflVal(v reflect.Value) {
	t := reflect.TypeOf(v.Interface())
	PrintWithReflType(t)
}

func PrintWithReflType(t reflect.Type) {
	//Print type statistics
	fmt.Print("\t\t\tType Statistics\n\n")
	fmt.Printf("%-20s | %-20s | \n%-20s | %-20d |", "Type", "Number of Fields", t.Name(), t.NumField())
	//Print type fields
	fmt.Print("\n\n\t\t\tType Fields\n\n")
	fmt.Printf("%-20s | %-20s | %-5s | Tags\n", "Field Name", "Type", "Index")
	fields := getFieldInfo(t)
	for _, field := range fields {
		printFieldInfo(t, field)
	}

	fmt.Println("------------------------------------------------------------------------------------------------------")
}

func printFieldInfo(t reflect.Type, field fieldInfo) {
	fmt.Printf("%-20s | %-20s | %-5d | ", field.Name, field.Type, field.Index)

	// Get all tags
	tag := t.Field(field.Index).Tag
	tagList := splitTags(tag)
	// Print tags
	for tagName, tagValue := range tagList {
		fmt.Printf("%s: %s | ", tagName, tagValue)
	}

	fmt.Println()
}

func splitTags(tag reflect.StructTag) map[string]string {
	tags := make(map[string]string)

	for _, tagName := range strings.Split(string(tag), " ") {
		if tagName == "" {
			continue
		}
		parts := strings.SplitN(tagName, ":", 2)
		key := parts[0]
		value := strings.Trim(parts[1], `"`)
		tags[key] = value
	}

	return tags
}

func getFieldInfo(t reflect.Type) []fieldInfo {
	var fields []fieldInfo
	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i).Type
		fieldName := t.Field(i).Name

		fields = append(fields, fieldInfo{
			Name:  fieldName,
			Type:  fieldType.Name(),
			Index: i,
		})

	}
	return fields
}
