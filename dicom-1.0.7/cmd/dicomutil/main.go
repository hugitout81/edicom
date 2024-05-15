package main

import (
	"flag"
	"fmt"
)

func main() {

	filepath := flag.String("path", "./", "a string")

	flag.Parse()

	fmt.Println("path = ", *filepath)

	files := get_dcm_files(*filepath)

	//verify each .dcm file
	for _, file := range files{	

		fmt.Println("File:", file)	
		validFile := verify(file)
		fmt.Println("Valid File?", validFile)
	}
}
