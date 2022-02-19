package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Good bye %v \n", n)
}

// 2 argument. calling the function for each element inside the array
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)

	}
}

// area of the circle
func cicleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// sayGreeting("Sergey")
	// sayBye("see later")
	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)

	a1 := cicleArea(10.5)
	a2 := cicleArea(15)
	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)

}
