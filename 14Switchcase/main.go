package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch in golang")

	rand.Seed(time.Now().UnixNano())

	diceNo := rand.Intn(6) + 1

	switch diceNo {
	case 1:
		fmt.Println("Dice number is 1 you can open")
		fallthrough // fallthrough- go to the next case also
	case 2:
		fmt.Println("Dice number is 2 you have 2 spots open")
	case 3:
		fmt.Println("Dice number is 3 you have 3 spots open")
	case 4:
		fmt.Println("Dice number is 4 you have 4 spots open")
	case 5:
		fmt.Println("Dice number is 5 you have 5 spots open")
	case 6:
		fmt.Println("Dice number is 6 spots you can roll up again ")
	default:
		fmt.Println("Wrong value !")
	}
}
