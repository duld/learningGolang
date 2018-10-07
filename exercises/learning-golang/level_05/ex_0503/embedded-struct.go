/*
Create a new type: vehicle.

The underlying type is a struct.
The fields:
  doors
  color

	Create two new types: truck & sedan.

The underlying type of each of these new types is a struct.
  Embed the “vehicle” type in both truck & sedan.
  Give truck the field “fourWheel” which will be set to bool.
  Give sedan the field “luxury” which will be set to bool. solution

Using the vehicle, truck, and sedan structs:
  using a composite literal, create a value of type truck and assign values to the fields;
  using a composite literal, create a value of type sedan and assign values to the fields.
Print out each of these values.

Print out a single field from each of these values.
*/
package main

import (
	"fmt"
)

// Vehicle : a type representing a generic car
type Vehicle struct {
	doors int
	color string
}

// Truck : represents a vehicle that is a truck
type Truck struct {
	Vehicle
	fourWheel bool
}

// Sedan : represents a vehicle that is a sedan
type Sedan struct {
	Vehicle
	luxury bool
}

func main() {
	myTruck := Truck{
		Vehicle: Vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	mySedan := Sedan{
		Vehicle: Vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println("myTruck:", myTruck)
	fmt.Println("mySedan:", mySedan)

	fmt.Println(myTruck.color)

	fmt.Println(mySedan.luxury)
}
