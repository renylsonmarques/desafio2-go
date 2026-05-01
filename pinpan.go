package main

import "fmt"

// Um jogo que pegue numeros de 1 a 100 e quando o numero for multiplo de 3 ele mostre "Pin"
// quando numero for multiplo por 5 mostre "Pan" e quando for multiplo de 3 e 5 mostre "PinPan"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
