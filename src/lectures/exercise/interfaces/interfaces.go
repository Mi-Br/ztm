//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import (
	"fmt"
)

type LiftSize int

const (
	Small Size = iota
	Medium
	Large
)

type LiftPicker interface {
	PickLift() int //should return integere corresponding to lift size
}

//will store vehicle name
type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) PickLift() int {
	return Small
}
func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}

func (c Car) PickLift() int {
	return Medium
}
func (c Car) String() string {
	return fmt.Sprintf("Car: %v", string(c))
}

func (t Truck) PickLift() int {
	return Large
}
func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

func directVehicles(l LiftPicker) {
	switch l.PickLift() { //switch on lift size
	case Large:
		fmt.Println("Directing to Large lift")
	case Medium:
		fmt.Println("Directing to Medium lift")
	case Small:
		fmt.Println("Directing to Small lift")
	}
}

func main() {

}
