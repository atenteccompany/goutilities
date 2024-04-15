package goutilities

import (
	"encoding/json"
	"fmt"
	"runtime"
)

// Logs any value on stdOut
func LogVal(msg string, val interface{}) {
	fmt.Printf("%-50s: %v\n", msg, val)
}

// Logs JSON representation of any objects on stdout
func LogJSON(val interface{}) {
	j, _ := json.MarshalIndent(val, "", "    ")
	fmt.Printf("%s\n", j)
}

// Logs heap allocation at time of call.
func LogHeapAlloc(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	msg = fmt.Sprintf("%s: Heap Alloc", msg)
	LogVal(msg, (m.HeapAlloc / 1024 / 1024))
}
