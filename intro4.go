package main

import "fmt"

func main() {
	//初期値　；　条件；　インクリメント等
	for index := 0; index <= 10; index++ {

		if index == 3 {
			continue
		} else if index == 8 {
			break
		}

		fmt.Println(index)
	}

	intarray := [...]int{0, 1, 2, 3, 4, 5}
	//len関数で配列の長さを取得できるらしい
	for i := 0; i < len(intarray); i++ {
		fmt.Println(intarray[i])
	}
}
