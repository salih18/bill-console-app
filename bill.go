package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{}, // {"cake": 2.99, "coffee": 1.99}
		tip:   0,
	}
	return b
}

func (b *bill) format() string {

	fs := "YOUR BILL \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}
	// list tip
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "Total:", total)

	return fs
}

// update tip
func (b *bill) updateTip(t float64) {
	b.tip = t
}

// add item
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill and write
func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		fmt.Println("Couldn't save the file")
	}
	fmt.Println("The bill has been saved into folder")
}
