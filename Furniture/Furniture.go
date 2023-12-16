package Furniture

import "fmt"

type Furniture struct {
	Name     string
	Length   int
	Height   int
	Width    int
	Depth    int
	Weight   float32
	Colour   string
	Material string
	Shape    string
}

type Furnitures struct {
	Furnitures []Furniture
}

type WallCabinet struct {
	Furniture Furniture
	Height    int
	Width     int
	Depth     int
	Weight    float32
	Colour    string
	Material  string
}
type FloorCabinet struct {
	Furniture Furniture
	Height    int
	Width     int
	Depth     int
	Weight    float32
	Colour    string
	Material  string
}
type Table struct {
	Furniture Furniture
	Height    int
	Width     int
	Weight    float32
	Colour    string
	Material  string
}
type Chair struct {
	Furniture Furniture
	Height    int
	Weight    float32
	Colour    string
	Material  string
}
type Chandelier struct {
	Furniture Furniture
	Height    int
	Colour    string
	Material  string
	Shape     string
}
type Bath struct {
	Furniture Furniture
	Length    int
	Height    int
	Depth     int
	Weight    float32
	Colour    string
	Material  string
}
type Sink struct {
	Furniture Furniture
	Length    int
	Height    int
	Depth     int
	Weight    float32
	Colour    string
	Material  string
	Shape     string
}
type Toilet struct {
	Furniture Furniture
	Height    int
	Weight    float32
	Colour    string
	Material  string
}
type Dresser struct {
	Furniture Furniture
	Length    int
	Height    int
	Depth     int
	Weight    float32
	Colour    string
	Material  string
}
type Sofa struct {
	Furniture Furniture
	Length    int
	Height    int
	Depth     int
	Weight    float32
	Colour    string
	Material  string
}
type Closet struct { //комод
	Furniture Furniture
	Height    int
	Width     int
	Depth     int
	Weight    float32
	Colour    string
	Material  string
}
type Armchair struct {
	Furniture Furniture
	Height    int
	Width     int
	Weight    float32
	Colour    string
	Material  string
	Shape     string
}

func WriteFirnuture() Furnitures {
	fmt.Println("Вся мебель имеющаяся в доме:")
	var furniture []Furniture
	furniture = append(furniture, Furniture{
		Name:     "Table",
		Height:   75,
		Width:    80,
		Depth:    0,
		Weight:   13.1,
		Colour:   "white",
		Material: "wood",
		Shape:    "square"})
	furniture = append(furniture, Furniture{
		Name:     "Chair",
		Height:   83,
		Width:    45,
		Depth:    0,
		Weight:   3.6,
		Colour:   "black",
		Material: "plastic",
		Shape:    "rectangular",
	})
	furniture = append(furniture, Furniture{
		Name:     "Chandelier",
		Height:   65,
		Width:    48,
		Depth:    48,
		Weight:   2.2,
		Colour:   "white",
		Material: "glass",
		Shape:    "author's",
	})
	furniture = append(furniture, Furniture{
		Name:     "Bath",
		Length:   150,
		Height:   59,
		Width:    70,
		Depth:    41,
		Weight:   16.5,
		Colour:   "white",
		Material: "acrylic",
		Shape:    "rectangular",
	})
	furniture = append(furniture, Furniture{
		Name:     "Sink",
		Length:   100,
		Height:   16,
		Depth:    50,
		Weight:   22.39,
		Colour:   "white",
		Material: "marble",
		Shape:    "rectangular",
	})
	furniture = append(furniture, Furniture{
		Name:     "WallCabinet",
		Height:   70,
		Width:    50,
		Depth:    19,
		Weight:   13,
		Colour:   "white",
		Material: "wood",
		Shape:    "rectangular",
	})
	furniture = append(furniture, Furniture{
		Name:     "Toilet",
		Height:   78,
		Width:    35,
		Depth:    62,
		Weight:   28,
		Colour:   "white",
		Material: "ceramic",
	})
	furniture = append(furniture, Furniture{
		Name:     "Dresser",
		Height:   80,
		Width:    47,
		Depth:    40,
		Weight:   28,
		Colour:   "white",
		Material: "wood",
		Shape:    "rectangular",
	})
	furniture = append(furniture, Furniture{
		Name:     "Sofa",
		Height:   70,
		Width:    130,
		Depth:    50,
		Weight:   45,
		Colour:   "black",
		Material: "flock",
		Shape:    "rectangular",
	})
	furniture = append(furniture, Furniture{
		Name:     "Closet",
		Height:   185,
		Width:    100,
		Depth:    70,
		Weight:   35,
		Colour:   "brown",
		Material: "wood",
		Shape:    "rectangular",
	})
	furniture = append(furniture, Furniture{
		Name:     "Armchair",
		Height:   83,
		Width:    45,
		Depth:    0,
		Weight:   9,
		Colour:   "black",
		Material: "flock",
	})
	fmt.Println(furniture)
	return Furnitures{Furnitures: furniture}
}
