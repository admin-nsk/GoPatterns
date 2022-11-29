package Creational

import "fmt"

type SingletonCon struct {
	url  string
	port string
}

var singletonConInstance *SingletonCon

func GetInstance() *SingletonCon {
	if singletonConInstance == nil {

		fmt.Println("Creating singleton instance")

		singletonConInstance = &SingletonCon{
			url:  "the url",
			port: "the port",
		}
	} else {
		fmt.Println("Singleton instance already created")
	}

	return singletonConInstance
}

func singletonMain() {
	for i := 0; i < 10; i++ {
		GetInstance()
	}
}
