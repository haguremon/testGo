package main

import "fmt"

//関数の種類多すぎわろた まだまだあるｗ
func sayhello() {

	fmt.Println("アーセナルの試合があるから寝る　Sleep")
}

func cal(x, y string) {

	fmt.Println(x + y)

}
func modorititairyou() (int, string, bool) {

	return 1, "aaa", true
}
func modoritinosengenn() (r string) {

	r = "aaaa"

	return
}

func main() {

	mumeikannsu := func() string {
		fmt.Println("aaa")
		return "wao"
	}
	mumeikannsu()

	sayhello()

	func(a string) {
		fmt.Println(a)
	}("aiueo")

	a, i, u := modorititairyou()

	fmt.Println(a, i, u)

	cal("SOLID ", "原則")

	modoritinosengenn()

}
