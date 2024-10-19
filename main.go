package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	var age int = 30
	var name string = "Krishna"
	var height float32 = 5.10
	fmt.Printf("My name is %s and my age is %d\n",name, age)
	fmt.Printf("My height is %.2f\n", height)

	place := "Erode"
	runs := 80
	fmt.Println("I am from", place)

	if runs > 100 {
		fmt.Println("You have scored a century")
	} else {
		fmt.Println("You have scored less than a century")
	}

	// i is only available inside the for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 10
	for i < 15 {
		fmt.Println(i)
		i++
	}

	j := 20
	for {
		if j > 25 { break }
		fmt.Println("Infinite loop")
		j++
	}

	
	nums := [5]int{ 100, 200, 300, 400, 500 }
	for index, value := range nums {
		fmt.Println(index, value)
	}

	// Variable declaration with data type says how much memory to allocate for the variable
	// For int64, 8 bytes of memory is allocated
	// For int32, 4 bytes of memory is allocated
	var x int64 = 10
	var y int64 = 20
	fmt.Println("Sum of the numbers is", add(x, y))
	fmt.Println("The value of x is", x)

	// Pointers 
	// xp can store address of int64 type variable
	// &y is the address of y
	var xp *int64
	xp = &x 
	fmt.Println("The value of xp is", xp)
	xp = &y 
	fmt.Println("The value of xp is", xp)


	// I want x to be 20 after add function
	// Pass by reference
	fmt.Println("Current Value of x is ", x)
	fmt.Println("Sum of the numbers is", addByReference(&x,y))
	fmt.Println("Value of x after add by reference is ", x)

	type Person struct {
		name string
		age int
		height float32
		place string
	}

	p := Person{ name: "Krishna", age: 30, height: 5.10, place: "Erode" }
	fmt.Println(p.name, p.age, p.height, p.place)

 

}

// Pass by value
func add(x int64, y int64) int64  {
	x = x + 10
	return x + y
}

// x has value 10
// &x has address value 0x14000090048
// *(&x)value in the address &x is 10
func addByReference(x *int64, y int64) int64 {
	*x = *x + 10
	return *x + y
}