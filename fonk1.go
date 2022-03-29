package main

import "fmt"

func main() {
	fmt.Println(add(5, 6))
	message := "hi"
	sayHello(&message)
	println(message)

}

func add(x, y int) int {
	f := x + y
	return f
}

func sayHello(message *string) { // pointer ... Adreste merhaba tutulduğundan main' de merhaba döner.
	println(message)
	*message = "Merhaba"
}
