package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup":          4.99,
		"pie":           7.99,
		"salad":         6.99,
		"toffee puddng": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key types
	phoneBook := map[int]string{
		26787648: "sergio",
		47590375: "mike",
		88500124: "mantana",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[26787648])

	phoneBook[47590375] = "John"
	fmt.Println(phoneBook)

	for k, v := range phoneBook {
		fmt.Println(k, "-", v)
	}

}
