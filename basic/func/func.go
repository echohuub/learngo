package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: " + op)
	}
}

// 13 / 3 = 4 ... 1
func div(a, b int) (q, r int) {
	//return a / b, a % b
	q = a / b
	r = a % b
	return
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Callint function %s with args "+
		"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func main() {
	fmt.Println(eval(3, 4, "/"))
	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 10, 2))

	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 10, 2))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
