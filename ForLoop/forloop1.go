package main

//import "fmt"
import (
	"os"
	"fmt"
	)

func main(){
	var result string
	var seperator string
	test :=  "hi"
	for i := 1; i < len(os.Args); i++ {
		result += seperator + os.Args[i]
		seperator = " "
	}
	fmt.Println(result)
	fmt.Println(test)
}
