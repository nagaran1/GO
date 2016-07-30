package main

import (
	"os"
	"fmt"
	"strings"
	)

func main(){
	var result string
	var seperator string
	for _, arg :=  range os.Args[1:] {
		result += seperator + arg
		seperator = " "
	}
	fmt.Println(result)

	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println(os.Args[1:len(os.Args)])
}
