package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("STARSHIP TAKE-OFF")

	g := rand.IntN(20) + 1
	w := rand.IntN(40) + 1
	r := g * w

	fmt.Printf("GRAVITY=%d\n", g)
	fmt.Println("TYPE IN FORCE")

	var f int

	for i := 0; i < 3; i++ {
		fmt.Scanln(&f)
		if f < r {
			fmt.Print("TOO LOW")
		}
		if f > r {
			fmt.Print("TOO HIGH")
		}
		if f == r {
			goto win
		}
		fmt.Println(", TRY AGAIN")
	}
	fmt.Println()
	fmt.Println("YOU FAILED -")
	fmt.Println("THE ALIENS GOT YOU")
	return
win:
	fmt.Println("GOOD TAKE OFF")
}
