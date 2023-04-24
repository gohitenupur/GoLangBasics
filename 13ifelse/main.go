package main

import "fmt"

func main() {
	fmt.Println("Golang ifelse")

	loginCount :=10
	var result string;

	if loginCount < 10{
		result = "less than 10"
	}else if loginCount>10 {
		result = "more than 10"
	}else{
	    result = "equal to 10"
    }

	fmt.Println(result)

	if num := 3;num %2 ==0 {
		fmt.Println("number is even",num)
	}else{
		fmt.Println("number is odd",num)
	}

   
}

