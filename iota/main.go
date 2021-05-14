package main

import "fmt"

const (
  a = iota
  b
  c
  d
)

func main()  {
	fmt.Println("iota test")
	fmt.Println(a,b,c,d)
}