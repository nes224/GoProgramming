package main 
import (
	"fmt"
)

type rectangle struct{
	lenght float64
	breadth float64
	color string

	geometry struct{
		area float64 
		perimeter float64
	}
}

func main(){
	var rect = &rectangle{}
	rect.lenght = 10
	rect.breadth = 20
	rect.color = "Green"

	rect.geometry.area = rect.lenght * rect.lenght
	rect.geometry.perimeter = 2 * (rect.lenght + rect.breadth)

	fmt.Println(rect)
	fmt.Println("Area:\t",rect.geometry.area)
	fmt.Println("Perimeter:",rect.geometry.perimeter)
}
