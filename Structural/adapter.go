package Structural

import "fmt"

type transport interface {
	navigateToDestination()
}

type client struct {
}

func (c *client) startNavigation(tr transport) {
	fmt.Println("Client starting the navigation process")
	tr.navigateToDestination()
}

type boat struct {
}

func (b *boat) navigateToDestination() {
	fmt.Println("Boat is navigating to island.")
}

type car struct {
}

func (c *car) driveToDestination() {
	fmt.Println(" Car is going to the destination")
}

//Adapter struct
type carAdapter struct {
	carTransport *car
}

func (ca *carAdapter) navigateToDestination() {
	fmt.Println("Adapter modify car to allow navigation")
	ca.carTransport.driveToDestination()
}

func start() {
	client1 := &client{}
	boat1 := &boat{}

	client1.startNavigation(boat1)
	car1 := &car{}
	carAdapter1 := &carAdapter{carTransport: car1}
	client1.startNavigation(carAdapter1)
}