package main

import "fmt"

func IsPrintable(s string) bool {
	compteur2 := 0
	compteur := 0
	for _, x := range s {
		compteur++
		if x >= 32 && x <= 126 {
			compteur2++
		}

	}
	if compteur == compteur2 {
		return true
	}
	return false
}

func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))
	fmt.Println(IsPrintable("Hello! How are you?"))

}
