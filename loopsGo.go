package main 
import "fmt"
func main(){
	nesArray := []string{"nes1","nes2","nes3"}
	for i:=0;i<4;i++{
		fmt.Println("Nes224")
	}
	for i,j := range nesArray{
		fmt.Println(i,j)
	}
	mmap := map[int]string{
		22:"k1",
		23:"k2",
		24:"k3",
	}
	// iterate loop key and value
	for key,value := range mmap{
		fmt.Println(key,value)
	}
}