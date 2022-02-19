package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) // change
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi", "mike", "peach", "browser"}
	sort.Strings(names)
	fmt.Println(names)

	// find := sort.SearchStrings(names, "mike")
	// fmt.Println(find)
	fmt.Println(sort.SearchStrings(names, "mike"))

}
