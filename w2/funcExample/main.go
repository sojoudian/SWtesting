package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	secondEcho("Hello World!")
	echo()
}

func add(a, b int) int {
	return a + b
}

func secondEcho(s string) string {
	return s
}

func echo() {
	fmt.Println("Echooo!")
}
