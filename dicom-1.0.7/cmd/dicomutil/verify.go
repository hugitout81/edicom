package main

import (
	"fmt"
	//"encoding/json"
	//"github.com/suyashkumar/dicom"
	//"io/ioutil"
	"os"
	//"strings"
	//"unicode"
)

func check(e error){
	if e!= nil{
		panic(e)
	}
}

func verify(s string) bool {
	var arr [132]byte
	copy(arr[128:132], "DICM")
	fmt.Printf("array: %v\n", arr)

	slice := arr[:]

	//dataset, _ := dicom.ParseFile("test.dcm", nil)

	//fmt.Println(dataset)

	//j, _ := json.Marshal(dataset)
	
	//fmt.Println(j)

	//_ = ioutil.WriteFile("test.txt", j, 0644)

	f, err := os.Open(s)
	check(err)

	b1 := make([]byte, 132)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	
	//str := string(b1[:])

	//str1 := strings.Map(func(r rune) rune{
	//	if unicode.IsPrint(r) {
	//		return r
	//	}
	//	return -1
	//}, str)

	//fmt.Println("the string is:", str1)

	//fmt.Println("length of str:", len(str))
	//fmt.Println("length of DICM:", len("DICM"))
	
	if Equal(b1, slice) {
		fmt.Println("this IS a DICOM file.")
		return true
	} else {
		fmt.Println("NOT A DICOM FILE!")
		return false
	}
}

func Equal(slice1, slice2 []byte) bool {
    if len(slice1) != len(slice2) {
        return false
    }

    for i := range slice1 {
        if slice1[i] != slice2[i] {
            return false
        }
    }

    return true
}
