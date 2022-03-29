package main

import "fmt"

func main() {
	//fmt.Println(toplam(6))
	a, b := toplam(5)
	fmt.Println("x:", a, "y:", b)
}

func toplam(sum int) (x, y int) { // x ve y direk isimlendirildi .
	x = sum * 5
	y = sum - x
	return // return x , y yazılmadı.
}
