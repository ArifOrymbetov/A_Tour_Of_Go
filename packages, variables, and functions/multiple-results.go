package pvf

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func Multiple_results() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
