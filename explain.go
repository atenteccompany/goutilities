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

// ExplainVal logs details of a struct value.
// It takes a reflect.Value as input and prints information about its type and fields.
func ExplainVal(v reflect.Value) {
	if v.Kind() != reflect.Struct {
		fmt.Print(color.HiRedString("Passed Value is not Struct\n"))
		return
	}

	t := v.Type()
	if t == nil {
		fmt.Print(color.HiRedString("Passed Type is nil\n"))
		return
	}
	ExplainType(t, 0)
}

// ExplainType logs details of a struct type.
// It takes a reflect.Type and depth as input and prints information about the type and its fields recursively.
func ExplainType(t reflect.Type, depth int) {
	if t == nil {
		fmt.Print(color.HiRedString("Passed Type is nil\n"))
		return
	}

	if t.Kind() != reflect.Struct {
		fmt.Print(color.HiRedString("Passed Type is not Struct\n"))
		return
	}

	primitiveFieldsCount, structFieldsCount := numFieldsByKind(t)

	//Print type statistics
	fmt.Print(color.HiRedString("\t\t\tType Statistics\n\n"))
	fmt.Printf("%-20s | %-20s | %-20s | %-20s | %-20s | \n", color.HiGreenString("Struct Name"), color.HiGreenString("Depth"), color.HiGreenString("Total Fields"), color.HiGreenString("Primitive Fields"), color.HiGreenString("Struct Fields"))
	fmt.Printf("%-20s | %-11d | %-12d | %-16d | %-13d |\n", color.HiWhiteString(t.Name()), depth, t.NumField(), primitiveFieldsCount, structFieldsCount)

	//Print type fields
	fmt.Print(color.HiRedString("\n\n\t\t\tType Fields\n\n"))
	fmt.Printf("%-24s | %-24s | %-24s | %-24s\n", color.HiGreenString("Field Name"), color.HiGreenString("Type"), color.HiGreenString("Index"), color.HiGreenString("Tags"))
	fields := getFieldInfo(t)
	for i, field := range fields {

		printFieldInfo(t, field)
		fieldType := t.Field(i).Type

		if fieldType.Kind() == reflect.Struct {
			//Print seperator after fields
			fmt.Println(color.HiWhiteString(strings.Repeat("-", 120)))
			ExplainType(fieldType, depth+1)
		}
	}

	if depth == 0 {
		fmt.Println(color.HiWhiteString(strings.Repeat("-", 120)))
	}
}

func printFieldInfo(t reflect.Type, field fieldInfo) {
	fmt.Print(color.WhiteString("%-15s | %-15s | %-15d | ", field.Name, field.Type, field.Index))

	// Get all tags
	tag := t.Field(field.Index).Tag
	tagList := splitTags(tag)
	// Print tags
	for tagName, tagValue := range tagList {
		fmt.Printf("%5s: %5s | ", tagName, tagValue)
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

func numFieldsByKind(t reflect.Type) (int, int) {
	primitiveCount := 0
	structCount := 0
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type.Kind() != reflect.Struct {
			primitiveCount++
		} else {
			structCount++
		}
	}
	return primitiveCount, structCount
}
