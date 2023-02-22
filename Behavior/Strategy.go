package Behavior

import "fmt"

type outputter interface {
	sendImage()
	getDestination() string
}

type printer struct {
	destination string
}

func (p printer) sendImage() {
	fmt.Println("printing image")
}

func (p printer) getDestination() string {
	return p.destination
}

type display struct {
	destination string
}

func (d display) sendImage() {
	fmt.Println("drawing image in screen")
}

func (d display) getDestination() string {
	return d.destination
}

func draw(output outputter) {
	fmt.Println("Preparing destination: ", output.getDestination())
	output.sendImage()
}

func runStrategy() {
	p := printer{destination: "printer machine 007"}
	d := display{destination: "image viewer"}

	draw(p)
	fmt.Println("----------------------------------")
	draw(d)
}
