package main

import "fmt"
// 使用递归实现阶乘
func Factorial(x int) (result int){
	if (x>1){
		result = x*Factorial(x-1)
	}else
	{
		result = 1
	}
	return result
}

func main()  {
	var num int = 10
	res := Factorial(num)
	fmt.Printf("%d 的阶乘是 %d \n", num,res)
}