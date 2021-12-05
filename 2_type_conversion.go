package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// CONVERT FLOAT TO INT
	var f64 float64 = 3.2
	var x int = int(f64)

	fmt.Println("FLOAT TO INTEGER", f64, x)

	// CONVERT INT TO STRING
	var someInt int = 32
	var fromInt = strconv.Itoa(someInt)

	fmt.Println("INT TO STRING", fromInt, reflect.TypeOf(fromInt))

	// CONVERT STRING TO INT64
	var someString string = "222"
	var fromString, _ = strconv.ParseInt(someString, 10, 64) // will return int64
	var fromInt64 int32 = int32(fromString) // convert to int32

	fmt.Println("STRING TO INT64", reflect.TypeOf(fromString), fromString, reflect.TypeOf(fromInt64), fromInt64)

}
