package main

import (
	"fmt"
)

func main() {
	const stirngdesuyo string = "Good night" //定数
	fmt.Println(stirngdesuyo)
	var intdesu int
	intdesu = 22
	stringDesu := "Good night" //変数
	stringDesu = "Good night\n9月"
	fmt.Println(stringDesu + fmt.Sprint(intdesu))
}
