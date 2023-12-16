package Accessories

import (
	"fmt"
)

type Accessory struct {
	Name string
	Length   int
	Height   int
	Width    int
	Depth    int
	Weight   float32
	Colour   string
	Material string
	Shape    string
}

type Accessories struct {
	Accessories []Accessory
}

type Brush struct {
	Accessories Accessory
	Length      int
	Weight      float32
	Colour      string
	Material    string
}
type SoapBox struct {
	Accessories Accessory
	Length      int
	Weight      float32
	Colour      string
	Material    string
	Shape       string
}
type Organizer struct {
	Accessories Accessory
	Height      int
	Width       int
	Weight      float32
	Colour      string
	Material    string
}
type Shelf struct {
	Accessories Accessory
	Length      int
	Weight      float32
	Colour      string
	Material    string
}
func WriteAccessories() Accessories {
	fmt.Println("Все аксессуары имеющиеся в доме:")
	var accessory []Accessory
	accessory = append(accessory, Accessory{
		Name: "Brush",
		Height:   36,
		Width:    8,
		Depth:    8,
		Weight:   0.09,
		Colour:   "white",
		Material: "steel",
	})
	accessory = append(accessory, Accessory{
		Name: "SoapBox",
		Height:   2,
		Width:    13,
		Depth:    10,
		Weight:   0.144,
		Colour:   "black",
		Material: "polyresin",
	})
	accessory = append(accessory, Accessory{
		Name: "Organizer",
		Height:   8,
		Width:    22,
		Depth:    22,
		Weight:   0.3,
		Colour:   "transparent",
		Material: "acrylic",
	})
	accessory = append(accessory, Accessory{
		Name:"Shelf" ,
		Height:   58,
		Width:    23,
		Depth:    23,
		Weight:   1.405,
		Colour:   "silver",
		Material: "steel",
	})
	fmt.Println(accessory)
	return Accessories{Accessories: accessory}
	}
