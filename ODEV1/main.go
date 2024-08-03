package main

import "fmt"

func main() {
	var islem int
	bakiye := 1000
	fmt.Printf("Merhabalar hesabınızda şuan %d bulunmaktadır. Lütfen para çekmek için (1)'i para yatırmak için (2)'yi tuşlayınız. ", bakiye)
	fmt.Scan(&islem)

	switch islem {
	case 1:
		fmt.Println("Ne kadar para çekmek istiyorsunuz?")
		var cekilecekTutar int
		fmt.Scan(&cekilecekTutar)
		if cekilecekTutar <= bakiye && cekilecekTutar > 0 {
			bakiye = bakiye - cekilecekTutar
			fmt.Printf("Hesabınızda %d TL bakiye kalmıştır", bakiye)
		} else {
			fmt.Println("Bu işlem yapılamaz.")
		}
	case 2:
		fmt.Println("Ne kadar para yatırmak istiyorsunuz?")
		var yatirilanTutar int
		fmt.Scan(&yatirilanTutar)
		bakiye = bakiye + yatirilanTutar
		fmt.Printf("Hesabınızda %d TL bakiye bulunmaktadır.", bakiye)
	default:
		fmt.Println("Böyle bir işlem bulunmamaktadır.")
	}
}
