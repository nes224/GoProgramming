package main

import "fmt"

func main(){
	fmt.Println(calculation("Rectangle",20,30))
	fmt.Println(calculation("Square",20))
	variadicExample(1,"red",true,10.5, []string{"foo","bar","baz"},map[string]int{"apple":23,"tomato":13})
}

func calculation(str string,y ... int) int{
	area :=1
	for _, val := range y{
		if str == "Rectangle"{
			area *= val
		}else if str == "Square"{
			area = val * val
		}
	}
	return area
}

func variadicExample(i ...interface{}){
	for _, v := range i{
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}