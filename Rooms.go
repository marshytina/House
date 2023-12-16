package House

import(
	"GolandProjects/awesome/House/Furniture"
	"GolandProjects/awesome/House/Appliances"
	"GolandProjects/awesome/House/Accessories"
	"GolandProjects/awesome/House/Decor"
	"GolandProjects/awesome/House/Dish"
)

)
type Room struct{
	Name string
}

type Kitchen struct {
	Room Room
	Table         Furniture.Table
	Chair         Furniture.Chair
	Chandelier    Furniture.Chandelier
	Dishwasher    Appliances.Dishwasher
	Kettle        Appliances.Kettle
	Coffeemachine Appliances.Coffeemachine
	ElectricStove Appliances.ElectricStove
	Fridge        Appliances.Fridge
	Mixer         Appliances.Mixer
	Multicooker   Appliances.Multicooker
	Pan           Dish.Pan
	Mug           Dish.Mug
	Pot           Dish.Pot
	FlatPlate     Dish.FlatPlate
	SoupPlate     Dish.SoupPlate
	Wineglass     Dish.Wineglass
	Fork          Dish.Fork
	Spoon         Dish.Spoon
	DessertSpoon  Dish.DessertSpoon
	Knife         Dish.Knife
	Spatula       Dish.Spatula
	CuttingBoard  Dish.CuttingBoard
}
type Bathroom struct {
	Room Room
	Bath               Furniture.Bath
	Sink               Furniture.Sink
	Toilet             Furniture.Toilet
	WallCabinet        Furniture.WallCabinet
	Dresser            Furniture.Dresser
	WashingMachine     Appliances.WashingMachine
	ElectricShaver     Appliances.ElectricShaver
	ElectricToothbrush Appliances.ElectricToothbrush
	Hairdryer          Appliances.Hairdryer
	CurlingIron        Appliances.CurlingIron
	Brush              Accessories.Brush
	SoapBox            Accessories.AccSoapBox
	Organizer          Accessories.Organizer
	Shelf              Accessories.Shelf
}
type Livingroom struct {
	Room Room
	Table          Furniture.Table
	Chair          Furniture.Chair
	Sofa           Furniture.Sofa
	Dresser        Furniture.Dresser
	Closet         Furniture.Closet
	Armchair       Furniture.Armchair
	FloorCabinet   Furniture.FloorCabinet
	TV             Appliances.TV
	PS5            Appliances.PS5
	AirConditioner Appliances.AirConditioner
	MusicCentre    Appliances.MusicCentre
	Laptop         Appliances.Laptop
	FloorLamp      Decor.FloorLamp
	Carpet         Decor.Carpet
	Bookcase       Decor.Bookcase
	Poof           Decor.Poof
	Pillow         Decor.Pillow
}
