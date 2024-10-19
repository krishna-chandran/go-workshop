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


}