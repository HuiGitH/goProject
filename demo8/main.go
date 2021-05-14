package main

import "fmt"

func main()  {
	const (
		a = iota
		b = iota
		c = iota
	)

	const (
		q =iota
		w
		e
	)
	fmt.Println(a, b ,c)
	fmt.Println(q,w,e)
}