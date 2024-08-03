package main

import (
	"fmt"
	"math/rand"
)

func main() {

	sayi := rand.Intn(100)
	var tahmin int
	fmt.Print("Tahmininizi giriniz : ")
	fmt.Scanln(&tahmin)
	for tahmin != sayi {
		if tahmin < sayi {
			fmt.Println("Daha büyük")
			fmt.Print("Tahmininizi giriniz : ")
			fmt.Scanln(&tahmin)
		}
		if tahmin > sayi {
			fmt.Println("Daha küçük")
			fmt.Print("Tahmininizi giriniz : ")
			fmt.Scanln(&tahmin)
		}
	}
	fmt.Printf("Tebrikler! Rastgele belirlenen sayimiz = %d", sayi)
}
