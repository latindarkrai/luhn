package car

import "fmt"

// Declares the Car struct to house the concept of a car
type Car struct {
	Name       string
	Doors      int
	CoolNumber int
}

// func (c Car) String() string {
//  	return fmt.Sprintf("Car: {Name: %s}", c.Name)
// }

func (c Car) String() string {
	return fmt.Sprintf("Car: {Name: %s, Doors: %d}", c.Name, c.Doors)
}

// This function is the designated initializer for the Car
func NewCar(carName string, doors int) Car {
	car := Car{Name: carName, Doors: doors}
	car.CoolNumber = 4 * car.Doors
	return car
}
