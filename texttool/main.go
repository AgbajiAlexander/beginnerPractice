package main

import (
	"fmt"
	"os"
	"strings"
)

//func main() {
//fmt.Println("I am Building text tool")

// if len(os.Args) != 3 {
// 	fmt.Println("Usage: go run . input.txt output.txt")
// 	return
// }

//inputFile := os.Args[1]
//outputFile := os.Args[2]

//data, err := os.ReadFile("input.txt")
//if err != nil {
//	fmt.Println("error occured while reading file", err)
//	return
//}

//newData := string(data)
//fmt.Println(newData)
// err := os.WriteFile(outputFile, data, 0064)
// if
//}

func main() {
	//fmt.Println("I am Building text tool")

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("error occured while reading file", err)
		return
	}

	newData := strings.ToUpper(string(data))
	//fmt.Println(newData)

	err = os.WriteFile(outputFile, []byte(newData), 0644)
	if err != nil {
		fmt.Println("error writing to output file", err)
		return

	}

}
