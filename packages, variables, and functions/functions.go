package pvf

import "fmt"

func add_for_file_functions(x int, y int) int {
	return x + y
}

func Functions() {
	fmt.Println(add_for_file_functions(42, 13))
}
