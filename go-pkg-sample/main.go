package main

import (
	"fmt"
	"go_samples/go-pkg-sample/stringutils"
)

func main(){
	s:="fabio"
	fmt.Println(stringutils.Upper(s))

	printMemoryAddr()
}

func printMemoryAddr(){
	var a = 10
	fmt.Printf("Address of a variable: %d\n", &a  )
}