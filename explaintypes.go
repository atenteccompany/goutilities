package goutilities

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/color"
)

type fieldInfo struct {
	Name  string
	Type  string
	Index int
}

func ExplainVal(v reflect.Value) {
	if v.Kind() == reflect.Ptr {
		fmt.Print(color.HiRedString("Passed Value is Pointer\n\n"))
		return
	}
	t := reflect.TypeOf(v.Interface())
	ExplainType(t)
}

func ExplainType(t reflect.Type) {
	if t == nil {
		fmt.Print(color.HiRedString("Passed Type is nil\n\n"))
		return
	} else if t.Kind() == reflect.Ptr {
		fmt.Print(color.HiRedString("Passed Type is Pointer\n\n"))
		return
	}
	//Print type statistics
	fmt.Print(color.HiRedString("\t\t\tType Statistics\n\n"))
	fmt.Printf("%-20s | %-20s | \n%-20s | %-17d|", color.HiGreenString("Type"), color.HiGreenString("Number of Fields"), color.HiWhiteString(t.Name()), t.NumField())
	//Print type fields
	fmt.Print(color.HiRedString("\n\n\t\t\tType Fields\n\n"))
	fmt.Printf("%-24s | %-24s | %-24s | %-24s\n", color.HiGreenString("Field Name"), color.HiGreenString("Type"), color.HiGreenString("Index"), color.HiGreenString("Tags"))
	fields := getFieldInfo(t)
	for _, field := range fields {
		printFieldInfo(t, field)
	}

	fmt.Println(color.HiWhiteString("------------------------------------------------------------------------------------------------------"))
}

func printFieldInfo(t reflect.Type, field fieldInfo) {
	fmt.Printf(color.WhiteString("%-15s | %-15s | %-15d | ", field.Name, field.Type, field.Index))

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
