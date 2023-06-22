package main 
import "fmt"

func area(l, w int)int{
     
    Ar := l* w
    return Ar
}

func main(){
	fmt.Println("Area of rectangle is : %.2f", area(12,10))
}