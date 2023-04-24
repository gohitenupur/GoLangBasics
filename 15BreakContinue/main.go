package main

import "fmt"

func main() {
	fmt.Println("Golang Break and continue")
    fmt.Println("--------------------------------")

	//create a slice
	// days := []string{"Sun", "Mon", "Tue", "Wed", "Thu","Fri"}

	// fmt.Println(days)

	// for i := 0; i < len(days); i++ {
    //     fmt.Println(days[i])
    // }

	// fmt.Println("--------------------------------")
	
	// i := 0
	// for i = range days{
	// 	fmt.Println(days[i])
	// }
	
    // fmt.Println("--------------------------------")

	// for i,val := range days{
	// 	fmt.Printf("Index is %v and value is %v \n ", i, val)
       
	// }

	itr :=1
	for itr < 10{
		// if itr ==5{
		// 	break;
		// }
		if itr == 2{
			goto val
		}

		if itr ==5 {
			itr++
			continue
		}
		fmt.Println("valie is",itr)
		itr++;
	}

	val :
	    fmt.Println("this is jump statement")
}