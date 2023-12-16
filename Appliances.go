package Appliances

import "fmt"

type Appliances struct { //бытовая техника
	Length    int
	Height    int
	Width     int
	Depth     int
	Name      string
	Weight    float32
	Colour    string
	Guarantee int
	Country   string
}

type Appliance struct {
	Appliance []Appliances
}

type WashingMachine struct {
	Appliances Appliances
	Height     int
	Length     int
	Width      int
	Name       string
	Weight     float32
	Colour     string
	Guarantee  int
	Country    string
}
type ElectricToothbrush struct {
	Appliances Appliances
	Length     int
	Name       string
	Colour     string
	Guarantee  int
	Country    string
}
type Hairdryer struct {
	Appliances Appliances
	Length     int
	Name       string
	Guarantee  int
	Country    string
}
type ElectricShaver struct {
	Appliances Appliances
	Length     int
	Name       string
	Guarantee  int
	Country    string
}
type CurlingIron struct {
	Appliances Appliances
	Length     int
	Name       string
	Guarantee  int
	Country    string
}
type Dishwasher struct {
	Appliances Appliances
	Height     int
	Width      int
	Depth      int
	Name       string
	Weight     float32
	Colour     string
	Guarantee  int
	Country    string
}
type Kettle struct { //чайник
	Appliances Appliances
	Height     int
	Volume     int
	Name       string
	Colour     string
	Guarantee  int
	Country    string
}
type Coffeemachine struct {
	Appliances Appliances
	Height     int
	Width      int
	Name       string
	Weight     float32
	Colour     string
	Guarantee  int
	Country    string
}
type ElectricStove struct {
	Appliances Appliances
	Height     int
	Width      int
	Name       string
	Weight     float32
	Colour     string
	Guarantee  int
	Country    string
}
type Fridge struct {
	Appliances Appliances
	Height     int
	Width      int
	Depth      int
	Name       string
	Weight     float32
	Colour     string
	Guarantee  int
	Country    string
}
type Mixer struct {
	Appliances Appliances
	Height     int
	Width      int
	Depth      int
	Name       string
	Weight     float32
	Colour     string
	Guarantee  int
	Country    string
}
type Multicooker struct {
	Appliances Appliances
	Height     int
	Width      int
	Depth      int
	Name       string
	Weight     float32
	Colour     string
	Guarantee  int
	Country    string
}
type TV struct {
	Appliances Appliances
	Height     int
	Width      int
	Name       string
	Weight     float32
	Colour     string
	Guarantee  int
	Country    string
}
type PS5 struct {
	Appliances Appliances
	Height     int
	Width      int
	Weight     float32
	Colour     string
	Guarantee  int
}
type AirConditioner struct {
	Appliances Appliances
	Height     int
	Width      int
	Name       string
	Weight     float32
	Colour     string
	Guarantee  int
	Country    string
}
type MusicCentre struct {
	Appliances Appliances
	Height     int
	Width      int
	Name       string
	Weight     float32
	Colour     string
	Guarantee  int
	Country    string
}
type Laptop struct {
	Appliances Appliances
	Height     int
	Width      int
	Name       string
	Weight     float32
	Colour     string
	Guarantee  int
	Country    string
}

func PrintAppliances() Appliance {
	fmt.Println("Вся бытовая техника имеющаяся в доме:")
	var appliances []Appliances
	appliances = append(appliances, Appliances{
		Height:    85,
		Width:     59,
		Depth:     40,
		Name:      "WashingMachine",
		Weight:    58,
		Colour:    "white",
		Guarantee: 1,
		Country:   "China"})
	appliances = append(appliances, Appliances{
		Length:    22,
		Name:      "ElectricTothbrush",
		Weight:    0.208,
		Colour:    "black",
		Guarantee: 2,
		Country:   "Germany"})
	appliances = append(appliances, Appliances{
		Length:    35,
		Name:      "Hairdryer",
		Weight:    0.84,
		Colour:    "black",
		Guarantee: 2,
		Country:   "China"})
	appliances = append(appliances, Appliances{
		Length:    17,
		Name:      "ElectricShaver",
		Weight:    0.330,
		Colour:    "black",
		Guarantee: 2,
		Country:   "China"})
	appliances = append(appliances, Appliances{
		Length:    32,
		Name:      "CurlingIron",
		Weight:    0.5,
		Colour:    "black",
		Guarantee: 2,
		Country:   "China"})
	appliances = append(appliances, Appliances{
		Height:    71,
		Width:     122,
		Depth:     8,
		Name:      "TV",
		Weight:    14.68,
		Colour:    "black",
		Guarantee: 1,
		Country:   "China"})
	appliances = append(appliances, Appliances{
		Height:    39,
		Width:     16,
		Depth:     9,
		Name:      "PS5",
		Weight:    3.9,
		Colour:    "white",
		Guarantee: 1,
		Country:   "Japan"})
	appliances = append(appliances, Appliances{
		Height:    77,
		Width:     49,
		Depth:     29,
		Name:      "AirConditioner",
		Weight:    26,
		Colour:    "white",
		Guarantee: 1,
		Country:   "China"})
	appliances = append(appliances, Appliances{
		Height:    72,
		Width:     56,
		Depth:     37,
		Name:      "MusicCentre",
		Weight:    7,
		Colour:    "black",
		Guarantee: 1,
		Country:   "China",
	})
	appliances = append(appliances, Appliances{
		Height:    38,
		Width:     27,
		Depth:     15,
		Name:      "Laptop",
		Weight:    1.9,
		Colour:    "black",
		Guarantee: 1,
		Country:   "China",
	})
	appliances = append(appliances, Appliances{
		Height:    89,
		Width:     64,
		Depth:     65,
		Name:      "Dishwasher",
		Weight:    40.58,
		Colour:    "silver",
		Guarantee: 3,
		Country:   "Italy",
	})
	appliances = append(appliances, Appliances{
		Height:    25,
		Width:     21,
		Depth:     18,
		Name:      "Kettle",
		Weight:    1.25,
		Colour:    "black",
		Guarantee: 1,
		Country:   "Italy",
	})
	appliances = append(appliances, Appliances{
		Height:    37,
		Width:     25,
		Depth:     43,
		Name:      "Coffeemachine",
		Weight:    7.5,
		Colour:    "black",
		Guarantee: 1,
		Country:   "Netherlands",
	})
	appliances = append(appliances, Appliances{
		Height:    85,
		Width:     50,
		Depth:     59,
		Name:      "ElectricStove",
		Weight:    38.8,
		Colour:    "black",
		Guarantee: 1,
		Country:   "Czech",
	})
	appliances = append(appliances, Appliances{
		Height:    185,
		Width:     60,
		Depth:     64,
		Name:      "Fridge",
		Weight:    65.4,
		Colour:    "black",
		Guarantee: 1,
		Country:   "Italy",
	})
	appliances = append(appliances, Appliances{
		Height:    38,
		Width:     25,
		Depth:     36,
		Name:      "Mixer",
		Weight:    6.9,
		Colour:    "black",
		Guarantee: 1,
		Country:   "China",
	})
	appliances = append(appliances, Appliances{
		Height:    28,
		Width:     24,
		Depth:     37,
		Name:      "Multicooker",
		Weight:    3.4,
		Colour:    "black",
		Guarantee: 1,
		Country:   "China",
	})
	fmt.Println(appliances)
	return Appliance{Appliance: appliances}

}
