package House

import(
	"fmt"
	"GolandProjects/awesome/House/Furniture"
	"GolandProjects/awesome/House/Accessory"
	"GolandProjects/awesome/House/Decor"
	"GolandProjects/awesome/House/Dish"
	"GolandProjects/awesome/House/Family"
)

type Home struct{
	Furniture Furniture.Furnitures
	Accessories Accessory.Accessories
	Decor Decor.Decors
	Dishes Dish.Dishes
	People Family.Persons
}

func CreateHouse() Home{

	Furniture: Furniture.WriteFurniture(),
	Accessories: Accessory.WriteAccessories(),
	Decor: Decor.WriteDecor(),
	Dishes: Dish.WriteDishes()
	People: Family.WriteFamily()
	fmt.Println(Furniture)

	}
}
