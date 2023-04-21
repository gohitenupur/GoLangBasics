// time is a spacific pkg in go documentation
package main

import (
	"fmt"
	"time"
)

func main(){
	// fmt.Println("Time study in golang")

	presentTime :=time.Now()
	fmt.Println("Present Time is :",presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Mon"))

	createDate :=time.Date(2020 ,time.August,10,23,23,0,0,time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Mon"))

	
}