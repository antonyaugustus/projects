package  main

import "fmt"

func main()  {
	fmt.Println("Hello World")

	d1 := dealership{name:"A1 Auto", location:"Denver"}
	fmt.Println("Location " + d1.location)

	var d2 dealership
	d2.name="Xtreme Auto"
	d2.location="Houston"
	fmt.Println(d2.name + d2.location)

	var d3 dealership
	d3 = dealership{name:"Clean Auto", location:"Dallas"}
	fmt.Println(d3.name + d3.location)
}

type dealership struct {
	name string
	location string

}