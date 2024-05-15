package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func get_dcm_files(fp string) []string {
	
	var file_strings = []string {}
	files, err := ioutil.ReadDir(fp)
   		if err != nil {
      		fmt.Println("Error:", err)
	  		return nil
  		}
   for _, file := range files {
	   fmt.Println(file.Name())
	   if strings.HasSuffix(file.Name(), ".dcm") {
			file_strings = append(file_strings, file.Name())	
		}
  }
	return file_strings
}
