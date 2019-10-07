package main

import "fmt"

func fe() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	a := [...]int{1, 2}

	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	fe()

	fmt.Println("Hello, world")
}
