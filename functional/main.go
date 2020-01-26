package main

import (
	"bufio"
	"fmt"
	"io"
	"learngo/functional/fib"
)

//type intGen func() int
//
//func (g intGen) Read(p []byte) (n int, err error) {
//	next := g()
//	if next > 10000 {
//		 return 0, io.EOF
//	}
//	s := fmt.Sprintf("%d\n", next)
//	return strings.NewReader(s).Read(p)
//}

func printFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib.Fibonacci()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	printFileContent(f)
}
