package goutilities

import (
	"reflect"
	"testing"
)

func TestExplainType(t *testing.T) {
	type Inner struct {
		F1_Inner int `json:"f1_inner"`
	}

	type Outer struct {
		F1 int           `json:"f1_outer"`
		F2 string        `json:"f2_outer"`
		F3 map[int]Inner `json:"f3_outer"`
		F4 Inner         `json:"f4_outer"`
	}

	v := Outer{
		F1: 0,
		F2: "",
		F3: map[int]Inner{},
		F4: Inner{
			F1_Inner: 0,
		},
	}

	ExplainType(reflect.TypeOf(v))
}
