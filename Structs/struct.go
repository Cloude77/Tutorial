package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(promt string, r *bufio.Reader) (string, error) {
	fmt.Print(promt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

// input user

func createBill() bill {
	reader := bufio.NewReader(os.Stdin) // unput

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Create the bill - ", b.name)
	return b
}

func promtOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promtOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("item added - ", name, price)
		promtOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		fmt.Println(tip)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promtOptions(b)
		}
		b.updateTip(t)

		fmt.Println("tip added - ", tip)
		promtOptions(b)

	case "s":
		b.save()
		fmt.Println("you saved the file - ", b.name)
	default:
		fmt.Println("that was not valid option ...")
		promtOptions(b)

	}
}

func main() {
	mybill := createBill()
	promtOptions(mybill)

}
