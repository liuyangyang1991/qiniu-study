package main
import "log"
import "time"

func add(a string)  {
	log.Println("defer:    "+a)
}

func desc()(string){
	defer  add(time.Now().String())
	return time.Now().String()
}

func main()  {
	log.Println("return:   "+desc())
}
