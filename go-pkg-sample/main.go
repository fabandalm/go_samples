package main

import (
	"fmt"
	"go_samples/go-pkg-sample/stringutils"
	"go_samples/go-pkg-sample/calc"
)

func main(){
	s:="fabio"
	fmt.Println(stringutils.Upper(s))
	result := calc.Add(2,3)
	println(result)
	printMemoryAddr()
}

func printMemoryAddr(){
	var a = 10
	fmt.Printf("Address of a variable: %d\n", &a  )
}