package main

import (
	"fmt"
	"strings"
)

type Aim struct {
	Flag     bool
	Location []int
}

func QueryStr(str1, str2 string) /*([]map[string] Aim) */ {
	//var info []map[string] Aim

	for i := 0; i < len(str2); i++ {
		var locationslice []int
		//var aimmap map[string] Aim
		//aimmap = make(map[string] Aim)

		//char := string(str2[i])
		//flag := strings.Contains(str1,string(str2[i]))
		for j := 0; j < len(str1); j++ {
			if strings.Contains(string(str2[i]), string(str1[j])) {
				aim := j
				locationslice = append(locationslice, aim)
			}

		}
		fmt.Printf("the result of letter %s is:", string(str2[i]))
		fmt.Println(locationslice)
		//aimmap[char] = Aim{flag,locationslice}
		//info = append(info,aimmap)
	}
	//return info
}

func main() {
	var str1 string
	var str2 string
	fmt.Println("input the aim string:")
	fmt.Scanf("%s\n", &str1)
	fmt.Println("input the query string:")
	fmt.Scanf("%s\n", &str2)
	QueryStr(str1, str2)
	/*lyy := QueryStr(str1,str2)
	var  sh []string
	for i,_:=range lyy  {
		lyyy := lyy[i]

		for k,_:=range str2  {
			ch := string(str2[k])
			sh = append(sh,ch)
		}
		for j,_:=range sh {
			fmt.Println(lyyy[string(sh[j])])
		}

		//fmt.Printf("the result of letter %s is %s\n",ch,lyyy[ch])

	}*/
}
