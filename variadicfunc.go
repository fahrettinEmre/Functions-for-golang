package main

import "fmt"

// çoklu dönüş
func main() {
	fmt.Println(swap("dalga", "emre"))
	a, b := swap("emre", "dalga")
	fmt.Println(b, a)

	numTerms, sum := add(2, 3, 4, 10)
	fmt.Println("Eklenen sayı:", numTerms, "\nToplamı:", sum)

	sayHello("Benim", "adım", "Emre", ".")

	c := add2(1, 2, 3, 4, 5, 10)
	fmt.Println(c)

}

func swap(x, y string) (string, string) {
	return y, x
}

func add(terms ...int) (int, int) { // variadic function (belirsiz sayıda parametre alan functionlar)
	result := 0
	for _, v := range terms {
		result += v
	}
	return len(terms), result
}

func sayHello(message ...string) { //
	for _, message := range message {
		println(message)
	}

}

func add2(terms ...int) int {
	result := 0
	for _, v := range terms {
		result += v
	}
	return result
}
