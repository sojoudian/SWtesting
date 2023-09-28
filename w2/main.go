package main

import "fmt"

func main() {
	var num int = 2 //2.0
	fmt.Println(num)
	echo()

	s := Talk()
	fmt.Println(s) // If you want to make a function accessible from other
	//programs you have to make the first letter Uppercase like
	//fmt.Println()
	fmt.Println(Talk())
	fmt.Println(Floop())

}

func echo() {
	var strg string = "Class 214"
	strg = "Class"
	// clss_num := 214
	var clss_num float64 = 214.55
	msg := "Go lang Rocks!!"
	msg = "Golang Rocks!!!!!"
	//msg = 2
	fmt.Println(strg, clss_num, msg)
}
