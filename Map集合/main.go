package main

import "fmt"

func main()  {
	var countryCapitalMap map[string]string

	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["China"] = "BeiJIng"
	countryCapitalMap["India"] = "New Delhi"

	/*使用key 输出map值*/
	for country := range countryCapitalMap{
		fmt.Println("Capital of ",country," is ", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在*/
	capital , exist := countryCapitalMap["United States"]
	if (exist){
		fmt.Println("Capital of United States is ", capital)
	}else{
		fmt.Println("Capital of United States is not present")
	}

	fmt.Println("")
	fmt.Println("删除前的map")
	for country := range countryCapitalMap{
		fmt.Println(country, countryCapitalMap[country])
	}
	// 删除元素
	fmt.Println("删除 China Node")
	delete(countryCapitalMap, "China")
	fmt.Println("China Node has deleted")

	fmt.Println("删除后的map")
	for country := range countryCapitalMap{
		fmt.Println(country, countryCapitalMap[country])
	}
}

