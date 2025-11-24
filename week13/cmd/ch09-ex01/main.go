package main

import "fmt"

type Kilometers float64
type Meters float64
type Miles float64

func (km Kilometers) ToMiles() Miles {
	return Miles(km / 1.609)
}
func (m Meters) ToMiles() Miles {
	return Miles(m / 1609)
}

func main() {
	kmph := Kilometers(160.1)
	fmt.Printf("%0.2f km/h equals %0.2f mile/p\n", kmph, kmph.ToMiles())
	meters := Meters(160100)
	
	fmt.Printf("%0.2f meter equals %0.2f mile\n", meters, meters.ToMiles())
}