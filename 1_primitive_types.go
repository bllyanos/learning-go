package main

import "fmt"

func main() {
	// 1 BOOLEANS
	var flag bool // zero value = false
	var trueFlag = true

	fmt.Println("BOOLEANS", flag, trueFlag)

	// 2 NUMERIC TYPES
	// 2.1 INTEGER TYPES - all int zero value is 0
	// note: we can use _ as thousand separator

	// 2.1.1 SIGNED INTEGERS
	var int8Example int8 = -67               // -128 to 127
	var int16Example int16 = 5_641           // -32768 to 32767
	var int32Example int32 = 72_300          // –2147483648 to 2147483647
	var int64Example int64 = 581_040_000_000 // –9223372036854775808 to 9223372036854775807

	fmt.Println("SIGNED INTEGERS", int8Example, int16Example, int32Example, int64Example)

	// 2.1.2 UNSIGNED INTEGERS - no negative value
	var uint8Example uint8 = 23                           // 0 to 255
	var uint16Example uint16 = 65_000                     // 0 to 65536
	var uint32Example uint32 = 3_500_000_000              // 0 to 4294967295
	var uint64Example uint64 = 18_446_744_073_709_551_615 // 0 to 18446744073709551615
	fmt.Println("UNSIGNED INTEGERS", uint8Example, uint16Example, uint32Example, uint64Example)

	// 2.1.3 SPECIAL INTEGERS
	var youKnowMe byte = 255 // same as uint8 - more common
	var imFlexible int = 32  // could be 32bit or 64bit depends on CPU

	fmt.Println("SPECIAL INTEGERS", youKnowMe, imFlexible)

	// 2.2 FLOATING POINT TYPES
	// 1.401298464324817070923729583289916131280e-45 to 3.40282346638528859811704183484516925440e+38
	var float32Example float32 = 3.14
	// 4.940656458412465441765687928682213723651e-324 to 1.797693134862315708145274237317043567981e+308
	var float64Example float64 = 4_322_182.992761273

	fmt.Println("FLOATING POINTS", float32Example, float64Example)

	// 3 STRINGS and RUNES
	var myChar rune = 'h'                // rune is an alias for int32
	var iCanAssignHere int32 = 'b'       // same as rune, but we use rune for chars and int32 for numbers
	var imGroot string = "I am Groot !!" // seems familiar

	fmt.Println("STRINGS and RUNES", string(myChar), string(iCanAssignHere), imGroot)
}
