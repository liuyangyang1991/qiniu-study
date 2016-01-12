package main

import (
	"time"
	"sync"
	"fmt"
)

func main()  {
	var a int64 = 100000
	for n:=0;n<10 ;n++  {
		var l sync.Mutex
		go func(){
			//l.Lock()
			//defer l.Unlock()
			for  {
				l.Lock()
				if a>0{
					a--
				}else {
					l.Unlock() //不能跳过break之前的unlock
					break
				}
				l.Unlock()
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(a)
}
