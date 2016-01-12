package main

import (
	"fmt"
)

type Student struct {
	StuNum  int
	StuName string
	StuAge  int
}

func query(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func main() {
	var StudentDB map[int]Student
	StudentDB = make(map[int]Student)
	var num, age int
	var name string
	var flag int
	for {
		fmt.Println("Please input stunum,stuname,stuage:")
		fmt.Scanf("%d %s %d", &num, &name, &age)
		StudentDB[num] = Student{num, name, age}
		fmt.Println("countinue?(1/0)")
		fmt.Scanf("%d", &flag)
		if flag == 0 {
			/*fmt.Println("Please input stunum,stuname,stuage:")
			fmt.Scanf("%d %s %d",&num,&name,&age)
			StudentDB[num] = Student{num,name,age}*/
			break
		}
	}
	//fmt.Println(len(StudentDB))
	var keyslice []int
	//keyslice = make([]int)
	for k, _ := range StudentDB {
		keyslice = append(keyslice, k)
		//fmt.Printf("%d,%s,%d\n",StudentDB[k].StuNum,StudentDB[k].StuName,StudentDB[k].StuAge)
	}
	for _, v := range query(keyslice) {
		fmt.Println(StudentDB[v])
	}
}
