package Decor

import "fmt"

type Decor struct {
	Name     string
	Length   int
	Height   int
	Width    int
	Depth    int
	Weight   float32
	Colour   string
	Material string
	Shape    string
	Diameter int
}

type Decors struct {
	Decors []Decor
}

type FloorLamp struct {
	Decor    Decor
	Height   int
	Weight   float32
	Colour   string
	Material string
	Shape    string
	Diameter int
}
type Carpet struct {
	Decor    Decor
	Length   int
	Width    int
	Weight   float32
	Colour   string
	Material string
}
type Bookcase struct {
	Decor    Decor
	Length   int
	Height   int
	Width    int
	Weight   float32
	Colour   string
	Material string
}
type Poof struct {
	Decor    Decor
	Length   int
	Weight   float32
	Colour   string
	Material string
	Shape    string
}
type Pillow struct {
	Decor    Decor
	Length   int
	Height   int
	Colour   string
	Material string
	Shape    string
}

func WriteDecor() Decors {
	fmt.Println("Вecь декор имеющийся в доме:")
	var decor []Decor
	decor = append(decor, Decor{
		Name:     "FloorLamp",
		Height:   170,
		Width:    84,
		Depth:    84,
		Weight:   6.4,
		Colour:   "transparent",
		Material: "glass",
	})
	decor = append(decor, Decor{
		Name:     "Carpet",
		Length:   100,
		Width:    50,
		Depth:    0,
		Weight:   6.4,
		Colour:   "grey",
		Material: "pile",
	})
	decor = append(decor, Decor{
		Name:     "Bookcase",
		Length:   104,
		Width:    70,
		Depth:    31,
		Weight:   21.74,
		Colour:   "white",
		Material: "wood",
	})
	decor = append(decor, Decor{
		Name:     "Poof",
		Length:   46,
		Width:    25,
		Weight:   4.9,
		Colour:   "biege",
		Diameter: 40,
		Shape:    "circle",
	})
	decor = append(decor, Decor{
		Name:     "Pillow",
		Length:   35,
		Width:    35,
		Depth:    15,
		Weight:   0.129,
		Colour:   "silver",
		Material: "polyester",
		Shape:    "square",
	})
	fmt.Println(decor)
	return Decors{Decors: decor}
}
