package main

import (
	"fmt"
	"strings"
)

func main() {

	main3()
	main4()
	main5()
}

func main3() {

	fmt.Println("I am string data type ")
	name := "sumit"

	info := "hello everyone my name is " + name + " i am learning golang "
	fmt.Println(info)

}

func main4() {
	message := "sumit chougule"
	name := "sumit"
	fmt.Printf("%c", message[0])
	// formating use above output s
	//if not %c used then unicode will get

	fmt.Println(len(message))

	fmt.Println((strings.Compare(message, name)))

	findChar := "Golang pogramming lang"

	fmt.Println((strings.Contains(findChar, "Golang")))
}

func main5() {
	stringOne := "this is some lowercase words"
	fmt.Println(strings.ToUpper(stringOne))

	stringtwo := "THIS IS SOME UPPERCASE WORDS"
	fmt.Println(strings.ToLower(stringtwo))
}
