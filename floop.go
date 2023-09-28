package main

func Floop() int {

//go build floop.go talk.go main.go

	sum := 0
	for i := 1; i < 5; i++ {
		//sum = sum + i
		sum += 1
	}
	return sum
}
