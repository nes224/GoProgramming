package main 
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)
type Config struct {
	SecretKey string `json:"secret_key`
}

func main(){
	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(fmt.Sprintf("Error reading configuration file: %v",err))
	}

	var config Config 
	err = json.Unmarshal(configData, &config)
	if err != nil{
		panic(fmt.Sprintf("Error parsing configuration data: %v",err))
	}

	if config.SecretKey == ""{
		panic("Secret key is missing from configuration")
	}

	fmt.Println("Application started successfully.")
}