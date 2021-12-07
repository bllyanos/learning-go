package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// DECISION (if / else)
	var someConditions bool = 3 < 5
	if someConditions {
		fmt.Println("condition is true")
	} else {
		fmt.Println("condition is false")
	}

	if false {
		fmt.Println("never printed")
	}

	// if else-if else
	myNum := 3
	if myNum < 5 {
		fmt.Println("Less than 5")
	} else if myNum > 5 {
		fmt.Println("More than 5")
	} else {
		fmt.Println("Is 5")
	}

	// GO DOESN'T HAVE TERNARY OPERATOR

	// we can declare if scoped variable
	if n := 5; n < 10 {
		fmt.Println("we can access n here", n)
	} else {
		fmt.Println("or here", n)
	}

	// LOOPS
	// infinite loop
	for {
		fmt.Println("this will loop forever")

		break // prevent infinite loop
	}

	// for i loop
	for i := 0; i < 10; i++ {
		fmt.Println("I:", i)
	}

	// seed random
	rand.Seed(time.Now().UnixNano())

	// condition only (like while loop)
	foundEven := false
	for !foundEven {
		random := rand.Intn(100)
		if random%2 == 0 {
			foundEven = true
			fmt.Println("FOUND :", random)
		} else {
			fmt.Println("TRY AGAIN:", random)
		}
	}

	// for range (like for each)
	// slice
	myNums := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for index, value := range myNums {
		fmt.Println("index:", index, "value:", value)
	}

	// map
	ages := map[string]int{
		"Billy":   23,
		"Irwanto": 27,
	}
	for key, value := range ages {
		fmt.Println(key, "is", value, "years old.")
	}

	// strings
	greeting := "Hello gophers !"
	// we can use label

someLabel:
	for index, value := range greeting {
		strValue := string(value)
		if strValue == "s" {
			break someLabel
		}
		fmt.Println("Greeting index:", index, "=>", strValue)
	}

	// SWITCH
	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
			fallthrough // break is default behaviour of go switch, so we use this keyword to fallthrough to next case
		default:
			fmt.Println(word, "is a long word!")
		}
	}

	// empty switch, case is evaluating condition
loopLabel:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break loopLabel
		default:
			fmt.Println(i, "is boring")
		}
	}
}
