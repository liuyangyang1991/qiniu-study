package main

import (
	"fmt"
	"github.com/ikbear/golang-training/mathutil"
)

func main() {

	a, err := mathutil.Division(6, 2)

	if err != nil {
		fmt.Println("除数不能为 0")
		return
	}
	fmt.Println("Division is: ", a)
}
