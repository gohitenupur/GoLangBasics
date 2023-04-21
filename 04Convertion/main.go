package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please rate our pizza between 1 and 5 ")

	scanner := bufio.NewReader(os.Stdin)

	rating,_ :=scanner.ReadString('\n')

	fmt.Println("Thanks for rating :",rating)
	fmt.Printf("Type of rating  %T \n ",rating)

	nextRating, err := strconv.ParseFloat(strings.TrimSpace(rating),64)

	if err!= nil {
        fmt.Println(err)
    }else{

	fmt.Println("Next rating is :",nextRating+1)
	}

}