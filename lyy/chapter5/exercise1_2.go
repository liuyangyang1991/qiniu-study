package main

import (
	"time"
	"sync/atomic"
	"fmt"
)

func main()  {
	var a int64 = 100000
	var l  int64
	for n:=0;n<10 ;n++  {
		 atomic.LoadInt64(&l)
		go func(){
			for  {
				if a>0{
					a--
				}else {
					break
				}
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(a)
}

