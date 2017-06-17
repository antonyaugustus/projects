// Variadic functions
package main

import "fmt"

func main() {
	d1 := dealership{"A1 Auto", "Houston"}
	dealerName := createDealerFullName(d1.name, d1.city)
	fmt.Println(dealerName)

	//fmt.Println("Location " + d1.city)

	var d2 dealership
	d2.name="Xtreme Auto"
	d2.city="Houston"
	//fmt.Println(d2.name + d2.city)

	var d3 dealership
	d3 = dealership{name:"Clean Auto", city:"Dallas"}
	//fmt.Println(d3.name + d3.city)

	printDealers(d1.name, d2.name, d3.name)

	sold, name := calculateInvUtil(100, 175, d1)
	fmt.Println(name, sold)
}

type dealership struct {
	name string
	city string
}

func printDealers(dealers ...string) {
	for _, dealerName := range dealers {
		fmt.Println(dealerName)
	}

}

func createDealerFullName(s1 string, s2 string) string {
	return s1 + " of " + s2
}

func calculateInvUtil(remaining float32, start float32, dealer dealership) (percentSold float32, dealerName string) {
	dealerName = dealerName + " sold "
	percentSold = (1-remaining/start)*100
	return
}