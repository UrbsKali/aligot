package main

// import utils.go

import (
	"fmt"
	"urbskali/aligot/utils"
)

func main() {
	test := utils.Matrix{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(test)

	test2 := utils.Matrix{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(test2)

	fmt.Println(test.Multiply(&test2))
}
