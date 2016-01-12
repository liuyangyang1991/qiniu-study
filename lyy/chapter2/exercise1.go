package main

import (
	"fmt"
)

func LenOfStr(str string) {
	a := len(str)
	fmt.Printf("the length of %s is: %d\n", str, a)
}
func main() {
	str := "HelloGolang!"
	LenOfStr(str)
}
