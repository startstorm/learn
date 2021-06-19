package main

import "fmt"

func main() {
	test1()
	test2()
}

func test1(){
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func test2()  {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}



