package Family

import "fmt"

type Family struct {
	Name          string
	Surname       string
	Age           int
	MaritalStatus string
	Profession    string
	Hobby         string
	FavouriteToy  string
}

type Persons struct {
	Persons []Family
}

type Mom struct {
	Name          string
	Surname       string
	Age           int
	MaritalStatus string
	Profession    string
}
type Dad struct {
	Name          string
	Surname       string
	Age           int
	MaritalStatus string
	Profession    string
}
type Son struct {
	Name          string
	Surname       string
	Age           int
	MaritalStatus string
	Profession    string
}
type Daughter struct {
	Name          string
	Surname       string
	Age           int
	MaritalStatus string
	Profession    string
}
type Grandma struct {
	Name          string
	Surname       string
	Age           int
	MaritalStatus string
	Hobby         string
}
type Grandpa struct {
	Name          string
	Surname       string
	Age           int
	MaritalStatus string
	Hobby         string
}


func WriteFamily() Persons {
	fmt.Println("Вся мебель имеющаяся в доме:")
	var family []Family
	family = append(family, Family{
		Name:          "Eva",
		Surname:       "Petrova",
		Age:           45,
		MaritalStatus: "Married",
		Profession:    "Teacher",
	})
	family = append(family, Family{
		Name:          "Ivan",
		Surname:       "Petrov",
		Age:           49,
		MaritalStatus: "Married",
		Profession:    "Doctor",
	})
	family = append(family, Family{
        Name:          "Oleg",
	    Surname:       "Petrov",
		Age:           23,
		MaritalStatus: "Married",
		Profession:    "Designer",
	})
	family = append(family, Family{
		Name:          "Veronika",
		Surname:       "Petrova",
		Age:           20,
		MaritalStatus: "Single",
		Profession:    "Singer",
	})
	family = append(family, Family{
		Name:          "Olga",
		Surname:       "Petrova",
		Age:           86,
		MaritalStatus: "Married",
		Hobby:         "Knitting",
	})
	family = append(family, Family{
		Name:          "Egor",
		Surname:       "Petrov",
		Age:           87,
		MaritalStatus: "Married",
		Hobby:         "Fishing",
	})
	fmt.Println(family)
	return Persons{Persons: family}
	}