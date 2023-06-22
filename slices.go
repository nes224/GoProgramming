package main

import (
	"fmt"
	"reflect"
)

func main(){
	var intSlice []int
	var strSlice []string 
	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Println(reflect.ValueOf(strSlice).Kind())

	employeeList := make(map[string]int)
	employeeList["Mark"] = 10
	employeeList["Sandy"] = 20
	fmt.Println(employeeList)

	fmt.Println(employeeList["Mark"])
	employeeList["Mark"] = 50
	fmt.Println(employeeList["Mark"])

	for k,v := range employeeList{
		fmt.Printf("key is %v and type is %v\n",k,v)
	}
}