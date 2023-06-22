package main 
import (
	"fmt"
	"reflect"
)

func main(){
	x := [...]int{10,20,30}
	fmt.Println(reflect.ValueOf(x).Kind())
	fmt.Println(len(x))

	for i :=0; i<len(x) ; i++{
		fmt.Println(x[i])
	}
}