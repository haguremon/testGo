package main

import (
	"fmt"
)

func main() {
	intarray := [...]int{0, 1, 2, 3}
	stringarray := [...]string{"a", "b", "c"}
	fmt.Println(stringarray)
	fmt.Println(intarray)
	fmt.Println(intarray[0],intarray[1],intarray[2],intarray[3])
}
