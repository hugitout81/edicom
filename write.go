package main

import (
	"fmt"
	"strings"
	"os"
)

func write_to_file(s string, filename string){

	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)
	
	if !strings.Contains(str, s) {
		str += s + "\n"
	}

	bytes := []byte(str)

	os.WriteFile(filename, bytes, 0644)
}
