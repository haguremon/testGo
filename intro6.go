package main

import "fmt"

type Animal struct {
	name, species, family string //型を一つだけでまとめて書くのみやすくて好きかも
	weight, tall          float64
}

func (r Animal) bmi(name string) {

	fmt.Println(r.name)
	fmt.Println("BMI:", r.weight/r.tall*2)

}
func (an Animal) sound(name, family, cry string) string {

	return name + "ha" + family + "de" + cry + "tonaku"
}

func main() {

	var neko = Animal{name: "neko", species: "Mammalia", family: "Felidae", weight: 4, tall: 47}

	fmt.Println(neko.name,
		neko.species,
		neko.family,
		neko.weight,
		neko.tall)

	neko2 := Animal{name: "neko2", species: "Mammalia", family: "Felidae", weight: 4, tall: 47}

	neko2.bmi(neko2.name)

	fmt.Println(neko2.sound("aaa", neko.family, "wowowo"))

}
