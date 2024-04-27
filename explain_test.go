package goutilities

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitTags(t *testing.T) {
	tests := []struct {
		name     string
		tag      reflect.StructTag
		expected map[string]string
	}{
		{
			name:     "TestSingleTag",
			tag:      `json:"name"`,
			expected: map[string]string{"json": "name"},
		},
		{
			name:     "TestMultipleTags",
			tag:      `json:"name" xml:"person"`,
			expected: map[string]string{"json": "name", "xml": "person"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, splitTags(test.tag))
		})
	}
}

func TestGetFieldInfo(t *testing.T) {
	type innerStruct struct {
		F1_1 string
	}

	type outerStruct struct {
		F1 int
		F2 string
		F3 struct{}
		F4 innerStruct
	}

	expected := []fieldInfo{
		{Name: "F1", Type: "int", Index: 0},
		{Name: "F2", Type: "string", Index: 1},
		{Name: "F3", Type: "", Index: 2},
		{Name: "F4", Type: "innerStruct", Index: 3},
	}

	t.Run("TestStructFields", func(t *testing.T) {
		fields := getFieldInfo(reflect.TypeOf(outerStruct{}))
		assert.Equal(t, expected, fields)
	})

}

func TestNumFieldsByKind(t *testing.T) {

	type innerStruct struct {
		F1_1 string
	}

	type outerStruct struct {
		F1 int
		F2 string
		F3 struct{}
		F4 innerStruct
	}

	t.Run("TestPrimitiveAndStructCounts", func(t *testing.T) {
		primitiveCount, structCount := numFieldsByKind(reflect.TypeOf(outerStruct{}))
		assert.Equal(t, 2, primitiveCount)
		assert.Equal(t, 2, structCount)
	})
}
