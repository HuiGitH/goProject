package main

import "fmt"

func main() {
   var a int = 10  
   var ip *int
   ip = &a
   fmt.Printf("变量的地址: %x\n", &a  )
   fmt.Printf("变量的指针地址: %x\n", ip)
   fmt.Printf("*ip变量的值: %d\n", *ip)

   var  ptr *int

   fmt.Printf("ptr 的值为 : %x\n", ptr  )
   fmt.Printf("ptr 的值为 : %v\n", ptr  )
   if (ptr ==nil){
	   fmt.Println("ptr是空指针")
   }
}