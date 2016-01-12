package main

import "fmt"

func ReverseStr(str string) {
	var str1 string
	for i := 0; i < len(str); i++ {
		var str2 string
		str2 = string(str[i])
		str1 = str2 + str1
	}
	fmt.Printf("the result of %s is %s\n", str, str1)
}

func main() {
	str := "HelloGolang!"
	ReverseStr(str)
}
