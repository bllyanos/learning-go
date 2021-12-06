package main

import "fmt"

func main() {
	// we can't change array length
	// declare an array with len 4
	// when we declare array with array literal less than its element, the rest element will have zero value
	var myArray [4]int = [4]int{1, 2, 3}

	fmt.Println(myArray)

	// length of array : len(arr)
	fmt.Println(len(myArray))

	// access 1 element
	fmt.Println("I'm zero :", myArray[3])
	fmt.Println("Element i2 :", myArray[2])

	// declare a slice
	// we can add element to a slice, because slice doesn't have fixed length
	var mySlice []int = []int{1, 2, 3, 4, 5, 6, 7}

	// create slice from array
	var sliceFromArr []int = myArray[:]

	fmt.Println(mySlice)
	fmt.Println("created from myArray", sliceFromArr)

	// when we change the value of an element of array slice, the original value (in array) also changed
	sliceFromArr[3] = 10
	fmt.Println("I changed the value to :", sliceFromArr[3], "-- test equality :", sliceFromArr[3] == myArray[3])

	// declaring slice with zero length, but reserve memory (capacity)
	// declaring slice with capacity help we append faster
	var myBigSlice []int = make([]int, 0, 10)
	fmt.Println(myBigSlice, len(myBigSlice), cap(myBigSlice))

	// append 3 (same as javascript's array.push)
	myBigSlice = append(myBigSlice, 3)

	fmt.Println("Now myBigSlice will have length", len(myBigSlice), "and capacity", cap(myBigSlice))
	fmt.Println(myBigSlice)

	// if we want to modify an element of slice without affecting original array, we can use copy
	var anotherArray = [5]int{5, 4, 3, 2, 1}
	var sliceFromThatArray = anotherArray[0:2] // will slice array from start (0 or we can leave it blank) to index 1
	fmt.Println(sliceFromThatArray)

	var copyOfThatSlice = make([]int, 2)
	copy(copyOfThatSlice, sliceFromThatArray)

	// we will change one of slice copy
	copyOfThatSlice[0] = 10

	fmt.Println("Tadaa...\n", "copy :", copyOfThatSlice, "vs", "original :", sliceFromThatArray)
}
