package goutilities

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/rodaine/table"
)

// ExplainType takes a struct type to log its details.
// It returns an error if the passed input is not a struct type.
func ExplainType(t reflect.Type) error {
	if t.Kind() != reflect.Struct {
		return errors.New("passed type must be a struct")
	}

	structQueue := make([]reflect.Type, 0)
	depthQueue := make([]int, 0)
	visitedTypes := make(map[string]bool)

	structQueue = append(structQueue, t)
	depthQueue = append(depthQueue, 0)

	for len(structQueue) > 0 {
		var structCnt int

		currentStruct := structQueue[0]
		currentDepth := depthQueue[0]
		structQueue = structQueue[1:]
        depthQueue = depthQueue[1:]
		structFields := make([]reflect.StructField, 0)

		for i := 0; i < currentStruct.NumField(); i++ {
			field := currentStruct.Field(i)

			structFields = append(structFields, field)

			first, second := fieldNewTypes(field)
			if first != nil && !visitedTypes[first.String()] {
				structQueue = append(structQueue, first)
				depthQueue = append(depthQueue, currentDepth+1)
				visitedTypes[first.String()] = true
			}
			if second != nil && !visitedTypes[second.String()] {
				structQueue = append(structQueue, second)
				depthQueue = append(depthQueue, currentDepth+1)
				visitedTypes[second.String()] = true
			}
			if first != nil || second != nil {
				structCnt++
			}
		}
		printStructHeader(currentStruct, structCnt, currentDepth)
		printStructFields(structFields)
	}

	return nil
}

func fieldNewTypes(field reflect.StructField) (reflect.Type, reflect.Type) {
	if field.Type.Kind() == reflect.Struct {
		return field.Type, nil
	}

	if field.Type.Kind() == reflect.Slice {
		return field.Type.Elem(), nil
	}

	if field.Type.Kind() == reflect.Map {
		var k, v reflect.Type
		if field.Type.Key().Kind() == reflect.Struct {
			k = field.Type.Key()
		}
		if field.Type.Elem().Kind() == reflect.Struct {
			v = field.Type.Elem()
		}
		if field.Type.Elem().Kind() == reflect.Slice {
			v = field.Type.Elem().Elem()
		}

		return k, v
	}

	return nil, nil
}

func printStructHeader(t reflect.Type, structCnt, depth int) {
	fmt.Printf("\033[1;32;4mStruct: %s(depth: %d, total fields count: %d, struct fields count: %d, primitive fields count: %d)\033[0m\n", t.Name(), depth, t.NumField(), structCnt, t.NumField()-structCnt)
}

func printStructFields(structFields []reflect.StructField) {
	tbl := table.New("Name", "Type", "Index", "Tags")
	tbl = tbl.WithHeaderSeparatorRow('-')

	for i, field := range structFields {
        ftype := field.Type.String()
        if field.Type.Kind() == reflect.Struct {
            ftype = "struct"
        }
		tbl.AddRow(field.Name, ftype, i, field.Tag)
	}

	tbl.Print()
	fmt.Println()
}
