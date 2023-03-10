package Behavior

import (
	"errors"
	"fmt"
)

type state interface {
	requestItem() error
	dispenseItem() error
}

type selectItemMachineState struct {
	vendingMachine *vendingMachine
}

func (si selectItemMachineState) requestItem() error {
	fmt.Println("Preparing selected item")
	si.vendingMachine.currentState = si.vendingMachine.itemRequest
	return nil
}

func (si selectItemMachineState) dispenseItem() error {
	return errors.New("first select an item")
}

type dispenseItemMachineState struct {
	vendingMachine *vendingMachine
}

func (di dispenseItemMachineState) requestItem() error {
	return errors.New("dispensing item, please wait")
}

func (di dispenseItemMachineState) dispenseItem() error {
	fmt.Println("Item dispensed")
	di.vendingMachine.currentState = di.vendingMachine.selectAvailable
	return nil
}

type vendingMachine struct {
	selectAvailable state
	itemRequest     state
	currentState    state
}

func (v vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func newVendingMachine() *vendingMachine {
	v := &vendingMachine{}
	selectItem := &selectItemMachineState{vendingMachine: v}

	dispenseItemState := &dispenseItemMachineState{vendingMachine: v}

	v.selectAvailable = selectItem
	v.itemRequest = dispenseItemState
	v.currentState = selectItem

	return v
}

func runState() {
	vm := newVendingMachine()

	fmt.Println("Please select an item")

	err := vm.requestItem()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("---")

	fmt.Println("Please select an item")

	err = vm.requestItem()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("---")
	err = vm.dispenseItem()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("---")

}
