package Structural

import "fmt"

// code related with computers
type computer interface {
	print()
	setPrinter(printer)
}

type desktop struct {
	printer printer
}

func (d *desktop) print() {
	fmt.Println("Print document for desktop pc")
	d.printer.printFile()
}
func (d *desktop) setPrinter(p printer) {
	d.printer = p
}

type laptop struct {
	printer printer
}

func (l *laptop) print() {
	fmt.Println("Print document for laptop")
	l.printer.printFile()
}

func (l *laptop) setPrinter(p printer) {
	l.printer = p
}

type printer interface {
	printFile()
}

type printerOne struct {
}

func (p printerOne) printFile() {
	fmt.Println("Printing by printer one")

}

type printerTwo struct {
}

func (p printerTwo) printFile() {
	fmt.Println("Printing by printer two")
}

func bridgeMain() {

	p1 := printerOne{}
	p2 := printerTwo{}

	desktopPC := desktop{}
	desktopPC.setPrinter(p1)
	desktopPC.print()
	fmt.Println()
	desktopPC.setPrinter(p2)
	desktopPC.print()
	fmt.Println()

	laptopPC := laptop{}
	laptopPC.setPrinter(p1)
	laptopPC.print()
	fmt.Println()
	laptopPC.setPrinter(p2)
	laptopPC.print()
	fmt.Println()

}
