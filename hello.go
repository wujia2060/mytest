package main

import "fmt"

func main() {
	a := []int{1,2,3,4,5}
	s1 := a[:3]
	s2 := a[3:]
	s3 := s2
	fmt.Println(s1,s2,s3)
}
xx
