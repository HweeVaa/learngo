package main

import (
	"fmt"
	"strings"
)

// func multiply(a, b int) int {
// 	return a * b
// }

func lenAndUpper(name string) (int, string){
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, upperName := lenAndUpper("Hwee")
	// totalLength, _ := lenAndUpper("Hwee") 밸류 무시, 4 출력
	fmt.Println(totalLength, upperName)
}