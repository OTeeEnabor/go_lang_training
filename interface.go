package main

import (
	"fmt"
)
//
//  declare interface
type motorVehichle interface {
	Mileage() float64
}

// 
type BMW struct {
	distance float64
	fuel float64
	averageSpeed string
}
type Audi struct {
	distance float64
	fuel float64
}

//  BMW method
func (b BMW) Mileage() float64{
	return b.distance/b.fuel
}
func (a Audi) Mileage() float64{
	return a.distance/a.fuel
}

func totalMileage(m []motorVehichle) {
	tm := 0.0
	for _, v := range m{
		tm = tm + v.Mileage()
	}
	fmt.Printf("Total mileage per month %f km/L", tm)


}
// type attendance interface {
// 	attendance()
// }

// type Bio struct {
// 	name   string
// 	school string
// }

// func (bio Bio) attendance() {
// 	fmt.Printf("%v was in school today \n", bio.name)
// }

func main() {
	b1 := BMW{
		distance: 167.9,
		fuel: 36,
		averageSpeed: "58",
	}

	a1 := Audi{
		distance: 152.9,
		fuel: 30,
	}

	person := []motorVehichle{b1, a1}
	totalMileage(person)


	// var person = Bio{
	// 	name:   "Edmund Rotimi",
	// 	school: "UJ",
	// }

	// checkAttendance(person)

}

// func checkAttendance(person attendance) {
// 	fmt.Printf("Hello")
// }
