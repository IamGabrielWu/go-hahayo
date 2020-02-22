package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// drill 01
func TestDrill01(t *testing.T) {
	for _, e := range os.Environ() {
		// fmt.Println(e)
		result := strings.Split(e, "=")
		// fmt.Println("key:  ", result[0], "value: ", result[1])
		if result[0] == "PATH" {
			arr := strings.Split(result[1], ":")
			for _, path := range arr {
				fmt.Println(path)
			}
		}

	}
}
