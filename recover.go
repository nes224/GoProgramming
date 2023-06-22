package main 

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Config struct{
	SecretKey string `json:"scret_key"`
}
func main(){
	defer func(){
		if r := recover(); r != nil {
			fmt.Printf("An error occurred: %v\n",r)
			fmt.Println("Application terminated gracefully.")
		}else{
			fmt.Println("Application executed successfully.")
		}
	}

	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(fmt.Sprintf("Error parsing configuration data: %v",err))
	}

	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		panic(fmt.Sprintf("Error parsing configuration data: %v",err))
	}

	if config.Secretkey == ""{
		panic("Secret key is missing from configuration")
	}

	fmt.Println("Application started succesfully.")
}