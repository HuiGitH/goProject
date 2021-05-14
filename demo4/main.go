package main

import (
	"fmt"
	"strconv"
)

//基本数据类型和string类型转换

func main()  {
	var num1 int =99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string // 空的string


	//方法一：使用fmt.Sprintf方法来转换成为字符串格式
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str = %q\n", str ,str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str = %q\n", str ,str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str = %q\n", str ,str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str = %q\n", str ,str)

	// 方法二：使用strconv包的函数

	var num3 int = 99
	var num4 float64 = 23.789
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str = %q\n", str,str)

	// strconv.FormatFloat(num4, 'f', 10, 64)
	// 'f'格式，  10：表示小数位保留10位， 64：表示这个小数是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str = %q\n",str ,str )

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str = %q\n",str, str)
	
	var num5 int64 = 4567
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T str = %q", str, str)
}