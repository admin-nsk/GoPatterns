package Structural

import "fmt"

type Driven interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car    Car
	driver *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 18 {
		c.car.Drive()
	} else {
		fmt.Println("Driver too young")
	}
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{Car{}, driver}
}

func proxy_main() {
	car := NewCarProxy(&Driver{12})
	car.Drive()
}
