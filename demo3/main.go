package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	var n1 = 100
	// fmt.Printf() 可以用于做格式化输出
	fmt.Printf("n1 的数据类型 %T", n1)

	//如何在程序查看某个变量的占用字节大小和数据类型(使用较多)
	var n2 int64 = 10

	fmt.Printf("n2的数据类型是 %T,n2占用的字节数是 %d\n", n2,unsafe.Sizeof(n2))

	var c1 byte = 'a'
	var c2 byte = 'b'
	var c3 = 'c'
	//当我们直接输出byte值时，就是输出了都赢得字符的码值
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	fmt.Println("c3=", c3)

	//如果我们希望输出对应字符，需要使用格式话输出
	fmt.Printf("c1 = %c c2 = %c, c3 = %c \n", c1, c2, c3)


}