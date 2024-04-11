package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// Format the bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// List items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// Add tip
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "tip:", b.tip)

	// Total
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "total:", total+b.tip)

	return fs
}

// Update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// Add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/normal/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("bill was saved to file")
}

func getInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	name, _ := getInput("Create a new bill name: ")

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func parsePrice(s string) float64 {
	var price float64

	_, err := fmt.Sscanf(s, "%f", &price)
	if err != nil {
		fmt.Println("the price must be a number")
		return 0
	}

	return price
}

func promptOptions(b bill) {
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ")

	switch opt {
	case "a":
		name, _ := getInput("Item name: ")
		price, _ := getInput("Item price: ")

		p := parsePrice(price)

		b.addItem(name, p)

		fmt.Println("item added -", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ")
		t := parsePrice(tip)

		b.updateTip(t)

		fmt.Println("tip added -", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("saving bill -", b.name)
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}
}
