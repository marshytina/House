package Dish

import "fmt"

type Dish struct {
	Length float32
	Height    int
	Name      string
	Weight    float32
	Colour    string
	Shape     string
	Material  string
	Volume    float32
	Diameter  float32
	Country   string
	Guarantee int
	Quantity  int
}
type Dishes struct {
	Dishes []Dish
}

type Pan struct {
	Dishes   Dish
	Weight   float32
	Colour   string
	Material string
	Volume   float32
	Diameter float32
	Quantity int
}
type Pot struct { //кастрюля
	Dishes   Dish
	Height   int
	Weight   float32
	Colour   string
	Material string
	Volume   float32
	Diameter float32
	Quantity int
}
type FlatPlate struct {
	Dishes   Dish
	Weight   float32
	Colour   string
	Material string
	Diameter float32
	Quantity int
}
type SoupPlate struct {
	Dishes   Dish
	Name     string
	Colour   string
	Material string
	Volume   float32
	Diameter float32
	Quantity int
}
type Mug struct { //кружка
	Dishes   Dish
	Height   int
	Name     string
	Weight   float32
	Colour   string
	Material string
	Volume   float32
	Quantity int
}
type Wineglass struct {
	Dishes   Dish
	Height   int
	Weight   float32
	Material string
	Volume   float32
	Quantity int
}
type Fork struct { //вилка
	Dishes   Dish
	Length   float32
	Weight   float32
	Colour   string
	Material string
	Quantity int
	Name     string
}
type Spoon struct {
	Dishes   Dish
	Length   float32
	Weight   float32
	Colour   string
	Material string
	Quantity int
	Name     string
}
type DessertSpoon struct {
	Dishes   Dish
	Length   float32
	Weight   float32
	Colour   string
	Material string
	Quantity int
	Name     string
}
type Knife struct {
	Dishes   Dish
	Length   float32
	Weight   float32
	Colour   string
	Material string
	Quantity int
	Name     string
}
type Spatula struct { //лопатка
	Dishes   Dish
	Length   float32
	Weight   float32
	Colour   string
	Material string
	Quantity int
	Name     string
}
type CuttingBoard struct {
	Dishes   Dish
	Length   float32
	Weight   float32
	Colour   string
	Material string
	Quantity int
	Name     string
}

func WriteDishes() Dishes {
	fmt.Println("Вся посуда имеющаяся в доме:")
	var dish []Dish
	dish = append(dish, Dish{
		Height:    27,
		Name:      "Pan",
		Weight:    0.52,
		Colour:    "black",
		Shape:     "round",
		Material:  "aluminium",
		Volume:    3,
		Diameter:  24,
		Country:   "China",
		Guarantee: 2,
		Quantity:  2,
	})
	dish = append(dish, Dish{
		Height:    30,
		Name:      "Pot",
		Weight:    2,
		Colour:    "silver",
		Shape:     "round",
		Material:  "steel",
		Volume:    4.5,
		Diameter:  24,
		Country:   "China",
		Guarantee: 2,
		Quantity:  2,
	})
	dish = append(dish, Dish{
		Name:     "FlatPlate",
		Weight:   0.35,
		Colour:   "black",
		Shape:    "round",
		Material: "porcelain",
		Diameter: 20,
		Quantity: 5,
	})
	dish = append(dish, Dish{
		Name:     "SoupPlate",
		Weight:   0.38,
		Colour:   "white",
		Shape:    "round",
		Material: "porcelain",
		Volume:   0.4,
		Diameter: 20,
		Quantity: 3,
	})
	dish = append(dish, Dish{
		Name:     "Mug",
		Weight:   0.34,
		Colour:   "white",
		Material: "porcelain",
		Volume:   0.25,
		Diameter: 12,
		Quantity: 4,
	})
	dish = append(dish, Dish{
		Name:     "Wineglass",
		Weight:   0.38,
		Colour:   "transparent",
		Shape:    "round",
		Material: "glass",
		Volume:   0.4,
		Diameter: 9,
		Quantity: 4,
	})
	dish = append(dish, Dish{ //вилка
		Length:   17.7,
		Weight:   20,
		Colour:   "silver",
		Material: "steel",
		Quantity: 10,
		Name:     "Fork",
	})
	dish = append(dish, Dish{
		Length:   17.5,
		Weight:   22.2,
		Colour:   "silver",
		Material: "steel",
		Quantity: 10,
		Name:     "Spoon",
	})
	dish = append(dish, Dish{
		Length:   13.5,
		Weight:   27,
		Colour:   "silver",
		Material: "steel",
		Quantity: 5,
		Name:     "DessertSpoon",
	})
	dish = append(dish, Dish{
		Length:   21.8,
		Weight:   68,
		Colour:   "silver",
		Material: "steel",
		Quantity: 10,
		Name:     "Knife",
	})
	dish = append(dish, Dish{
		Length:   25,
		Weight:   15,
		Colour:   "brown",
		Material: "wood",
		Quantity: 1,
		Name:     "Spatula",
	})
	dish = append(dish, Dish{
		Length:   36.8,
		Weight:   20,
		Colour:   "beige",
		Material: "bamboo",
		Quantity: 3,
		Name:     "CuttingBoard",
	})
	fmt.Println(dish)
	return Dishes{Dishes: dish}
	}
