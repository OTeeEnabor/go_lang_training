package main

import (
	"fmt"
	"math"
)
type Abser interface{
	Abs() float64
}
func main() {
	// declare variable a as Abser interface type
	var a Abser
	// declare variable f as type MyFloat and assign is - math.Sqrt2
	f := MyFloat(-math.Sqrt2)
	// declare variable v as type Vertex and assign it values 3,4
	v := Vertex{3,4}
	
	a = f // assign a the value of f
	fmt.Printf("%T, %v ",a,a)
	fmt.Println("")
	fmt.Printf("%v",v)
	fmt.Println("##########################")
	a = &v // assign a to the value at address v
	fmt.Printf("%T, %v ",a,a)
	fmt.Println("")
	fmt.Printf("%v",v)

	// a = v 
	fmt.Println(a.Abs())
}
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}













































//
// //  declare interface
// type motorVehichle interface {
// 	Mileage() float64
// }

// // 
// type BMW struct {
// 	distance float64
// 	fuel float64
// 	averageSpeed string
// } 
// // define a method for type bmw 
// type Audi struct {
// 	distance float64
// 	fuel float64
// }

// //  BMW method
// func (b BMW) Mileage() float64{
// 	return b.distance/b.fuel
// }
// func (a Audi) Mileage() float64{
// 	return a.distance/a.fuel
// }

// func totalMileage(m []motorVehichle) {
// 	tm := 0.0
// 	for _, v := range m{
// 		tm = tm + v.Mileage()
// 	}
// 	fmt.Printf("Total mileage per month %f km/L", tm)


// }
// // type attendance interface {
// // 	attendance()
// // }

// // type Bio struct {
// // 	name   string
// // 	school string
// // }

// // func (bio Bio) attendance() {
// // 	fmt.Printf("%v was in school today \n", bio.name)
// // }

// func main() {
// 	b1 := BMW{
// 		distance: 167.9,
// 		fuel: 36,
// 		averageSpeed: "58",
// 	}

// 	a1 := Audi{
// 		distance: 152.9,
// 		fuel: 30,
// 	}

// 	person := []motorVehichle{b1, a1}
// 	totalMileage(person)


