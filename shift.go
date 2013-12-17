package main

import (
	"fmt"
)

func main() {
	a := []bool{false, false, true, false, false}

	fmt.Println(a)
	fmt.Println(append(a[1:], false))
	fmt.Println(append([]bool{false}, a[:len(a)-1]...))
}
