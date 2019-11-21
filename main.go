package main

import "fmt"

func main()  {
	fmt.Println("Hello~module")
	a := []int{1,3,4,5}
	fmt.Println(a[:len(a)-1])
	fmt.Println(a[:0])
}
