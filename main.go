package main

import (
	"flag"
	"fmt"
	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
	"strings"
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

	for _, file := range files{	
	
		dataset, _ := dicom.ParseFile(file, nil)
	
		patId, _ := dataset.FindElementByTag(tag.PatientID)
		patName, _ := dataset.FindElementByTag(tag.PatientName)

		fmt.Println("Patient ID:", patId)
		fmt.Println("Name:", patName)
		
		strs := append(patId.Value.GetValue().([]string), patName.Value.GetValue().([]string)...)


		res := strings.Join(strs, "//") 		
		fmt.Println(res)
		write_to_file(res, "pat.txt")
		
	}
	


}
