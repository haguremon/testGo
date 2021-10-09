package main

import "fmt"

func main() {

	//簡略分めっちゃ便利やんけ
	if score := 49; score >= 90 {
		fmt.Println("goukau", fmt.Sprintln(score))
	} else if score > 50 {
		fmt.Println("gomidakara garbagecollection siyo")
	} else {
		fmt.Println("otu")
	}

	x := 128
	y := 128

	if binary := x + y; binary == 256 {
		fmt.Println("kirigaii")
	}
}
