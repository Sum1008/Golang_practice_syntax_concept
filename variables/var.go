package main

import "fmt"

var user = "Sumit"

func main() {
	var fruits string = "mango"

	fmt.Println(fruits)
	fmt.Println(user)
	fmt.Println("hello world")
	main2()
	main3()
	main4()
	main6()
	main7()
	main8()
}

func main2() {
	varone := 100
	varTwo := 2

	fmt.Println(varone)
	fmt.Println(varTwo)

}

func main3() {
	var name string = "sumit"
	var sirname = "chougule"

	number := 5

	fmt.Print(name, " ")
	fmt.Println(sirname)
	fmt.Println(number)
}

func main4() {

	var (
		num       int
		number    int    = 1
		greetings string = "hello"
	)
	fmt.Println(num)
	fmt.Println(number)
	fmt.Println(greetings)
}

func main6() {
	var one, two, three, four, five, six int = 1, 2, 3, 4, 5, 6

	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
	fmt.Println(four)
	fmt.Println(five)
	fmt.Println(six)

}

const user1 = "admin"

func main7() {

	fmt.Println("user1")
}

func main8() {

	isGolanfPL := true
	isHtmlPL := false

	fmt.Println(isGolanfPL)
	fmt.Println(isHtmlPL)
}
