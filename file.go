package goutilities

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bytedance/sonic"
)

// Writes JSON representation to file
func WriteJSONToFile(o interface{}) (string, error) {
	var b []byte
	var err error

	if b, err = sonic.Marshal(o); err != nil {
		return "", err
	}

	s := strings.Builder{}
	if _, err = s.Write(b); err != nil {
		return "", err
	}

	var fn string
	if fn, err = Writetofile(s.String()); err != nil {
		return "", err
	}

	return fn, nil
}

// Writes any string value to file.
// File will be created into directory output.
// File name will be prefixed with current timestamp.
func Writetofile(str string) (string, error) {
	strQueryTextFileName := fmt.Sprintf("output/output_%v.txt", time.Now().UnixMilli())
	file, err := os.Create(strQueryTextFileName)
	if err != nil {
		fmt.Printf("Failed to create file '%s': %s\n", strQueryTextFileName, err)
		return strQueryTextFileName, nil
	}
	defer file.Close()
	// ------------------------------------
	// Write the string select ....
	if _, err := file.WriteString(str); err != nil {
		fmt.Printf("Failed to write to file '%s': %s\n", strQueryTextFileName, err)
		return strQueryTextFileName, err
	}

	fmt.Printf("wrote to file '%s'\n", strQueryTextFileName)

	return strQueryTextFileName, nil
}
