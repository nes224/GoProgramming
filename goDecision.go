package main 
import "fmt"

func main(){
	v := 1200
	if v< 1200{
		fmt.Println(" v is less than 1000 \n")
	}else{
		fmt.Println("v is greater than 1000 \n")
	}

	v1 := 400
	v2 := 700
	v3 := 1200
	if (v1 == 400 && v2 == 700 && v3 >1200){
		fmt.Println("Value of v1 is 400 and v2 is 700 \n");
	}else{
		fmt.Println("Value of v3 is more than 1200")
	}
}