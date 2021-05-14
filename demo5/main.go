package main

import (
	"fmt"
	"strconv"
)


func main(){
	var str string = "true"
	var b bool
	var err_info error
	// b,_ = strconv.ParseBool(str)
	// 说明 
	// 1. strconv,ParseBool(str) 函数会返回两个值（value bool， err  error）
	
	b, err_info  = strconv.ParseBool(str)
	fmt.Printf("b type %T b= %v\n", b,b)
	fmt.Printf("err_info type %T b= %v\n", err_info,err_info)


	var str2 string  = "123456"
	var n1 int64
	var n2 int

	n1,_ = strconv.ParseInt(str2,10,64)
	n2 = int(n1)
	fmt.Printf("n1 type %T n1 = %v\n", n1,n1)
	fmt.Printf("n2 type %T n1 = %v\n", n2,n2)

	var str3 string = "123.456"
	var f1 float64

	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1 = %v\n", f1,f1)
	
}