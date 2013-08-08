package main

import (
	"fmt"
)

func the_first_function() {
	fmt.Println("the 1st function was called")
}

func the_second_function() {
	fmt.Println("the 2nd function was called")
}

func call_function(which_function func()) {
	which_function()
}

func main() {
	f1 := the_first_function
	call_function(f1)

	f2 := the_second_function
	f2()
}
