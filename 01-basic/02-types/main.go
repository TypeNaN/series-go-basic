package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	// boolean
	var boolA bool
	var boolB = true
	boolC := false
	fmt.Println("\n-- BOOL --")
	fmt.Println(boolA)
	fmt.Println(boolB)
	fmt.Println(boolC)

	// string
	var sA string
	var sB = "String B"
	sC := "String C"
	fmt.Println("\n -- STRING --")
	fmt.Println(sA)
	fmt.Println(sB)
	fmt.Println(sC)

	// Array
	var arrA [6]string
	var arrB = [6]string{"zero", "one", "two", "three", "four", "five"}
	arrC := [6]string{"zero", "one", "two", "three", "four", "five"}
	arrD := [6]string{"zero", "one", "two"}
	arrE := [6]string{2: "two", 3: "three"}
	fmt.Println("\n -- ARRAY --")
	fmt.Println(arrA)
	fmt.Println(arrB)
	fmt.Println(arrC)
	fmt.Println(arrD)
	fmt.Println(arrE)
	fmt.Println(arrB[0])
	fmt.Println(arrB[:3])
	fmt.Println(arrB[3:])
	fmt.Println(arrB[2:4])

	// Slice
	var slA []string
	var slB = []string{"zero", "one", "two", "three", "four", "five"}
	slC := []string{"zero", "one", "two", "three", "four", "five"}
	slD := []string{"zero", "one", "two", "three", "four", "five"}
	slE := []string{2: "two", 3: "three"}
	slF := make([]string, 3, 6)
	fmt.Println("\n -- SLICE --")
	fmt.Println(slA)
	fmt.Println(slB)
	fmt.Println(slC)
	fmt.Println(slD)
	fmt.Println(slE)
	fmt.Println(slF)
	fmt.Println(slB[0])
	fmt.Println(slB[:3])
	fmt.Println(slB[3:])
	fmt.Println(slB[2:4])

	slB = append(slB, "six")
	slB = append(slB, "seven")
	fmt.Println(slB)

	// int uint float
	var zero int
	var one = 1
	two := 2
	var f32 float32 = 0.32
	var f64 float64 = 0.46

	var c64 complex64 = complex(2, 3)
	var c128 complex128 = cmplx.Sqrt(-5 + 12i)

	fmt.Println("\n-- INT, UINT, FLOAT, COMPLAX --")
	fmt.Println(zero)
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(f32)
	fmt.Println(f64)
	fmt.Println(c64)
	fmt.Println(c128)

	var maxInt8 int8 = (1<<8 - 1) / 2
	var maxInt16 int16 = (1<<16 - 1) / 2
	var maxInt32 int32 = (1<<32 - 1) / 2 // type rune alias int32 represents a Unicode code point
	var maxInt64 int64 = (1<<64 - 1) / 2

	var hex uint8 = 0xFF
	var maxUInt8 uint8 = 1<<8 - 1 // type byte alias uint8
	var maxUInt16 uint16 = 1<<16 - 1
	var maxUInt32 uint32 = 1<<32 - 1 // type uintptr alias uint32
	var maxUInt64 uint64 = 1<<64 - 1

	fmt.Println("int   8 bit (-", maxInt8, ") ~ (+", maxInt8, ")")
	fmt.Println("int  16 bit (-", maxInt16, ") ~ (+", maxInt16, ")")
	fmt.Println("int  32 bit (-", maxInt32, ") ~ (+", maxInt32, ")")
	fmt.Println("int  64 bit (-", maxInt64, ") ~ (+", maxInt64, ")")

	fmt.Println("uint    hex (0 ~", hex, ")")
	fmt.Println("uint  8 bit (0 ~", maxUInt8, ")")
	fmt.Println("uint 16 bit (0 ~", maxUInt16, ")")
	fmt.Println("uint 32 bit (0 ~", maxUInt32, ")")
	fmt.Println("uint 64 bit (0 ~", maxUInt64, ")")
	
	const donotchange = "this is can't change"
	fmt.Println("\n-- CONST --")
	fmt.Println(donotchange)
	
	type Person struct {
		Name string
		Age int
	}

	var A = Person {
		Name : "MyName",
		Age : 16,
	}
	fmt.Println(A.Name)
	fmt.Println(A.Age)
}
