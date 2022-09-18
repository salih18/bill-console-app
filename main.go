package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput(reader, "Create a new bill name: ")

	b := newBill(name)
	return b

}

func getInput(r *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	name, err := r.ReadString('\n')
	name = strings.TrimSpace(name)
	return name, err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	choice, _ := getInput(reader, "Choose option (a - add item, s - save bill, t - add tip)")

	switch choice {
	case "a":
		item, _ := getInput(reader, "Item name: ")
		price, _ := getInput(reader, "Item price: ")
		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Wrong price")
			promptOptions(b)
		}
		b.addItem(item, p)
		promptOptions(b)
	case "s":
		b.save()
	case "t":
		tip, _ := getInput(reader, "Tip ammount: ")
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Tip must be number")
			promptOptions(b)
		}

		b.updateTip(t)
		promptOptions(b)
	default:
		fmt.Println("This is not a valid option")
		promptOptions(b)
	}

}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
