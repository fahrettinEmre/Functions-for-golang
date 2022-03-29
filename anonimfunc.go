package main

// anonim fonksiyonlar dışarıdan tekrar tekrar erişilemez ve tek bir kez kullanılırlar. işi biter ve hafızadan yok olur.
func main() {
	addFunc := func(terms ...int) (numTers int, sum int) { // buradaki func addFunc a eşitlenerek anonim hale gelmiştir.
		for _, v := range terms {
			sum += v
		}
		numTers = len(terms)
		return
	}

	numTerms, sum := addFunc(2, 5, 3)
	println("Eklenen:", numTerms, "Toplam:", sum)
}
