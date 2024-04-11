package main

import (
	"fmt"
	"os"
)

type shareBill struct {
	name            string
	numbersOfPeople int
	totalAmount     float64
	tip             float64
}

func newShareBill(name string) shareBill {
	sb := shareBill{
		name: name,
	}

	return sb
}

// Create new bill
func createShareBill() shareBill {
	var name string
	fmt.Println("What is the name of the share bill?")
	fmt.Scanln(&name)

	sb := newShareBill(name)
	fmt.Println("Created the share bill -", name)

	return sb
}

func saveShareBill(sb *shareBill) {
	data := []byte(sb.formatShareBill())

	err := os.WriteFile("bills/share/"+sb.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file")
}

// Format the bill
func (sb *shareBill) formatShareBill() string {
	fs := "Share bill breakdown: \n"

	// Total
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "total amount:", sb.totalAmount)

	// Tip
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "tip:", sb.tip)

	// Tip percentage
	fs += fmt.Sprintf("%-25v ...%0.2f%%\n", "tip percentage:", sb.calculateTipPercentage())

	// Total with tip
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "total with tip:", sb.calculateTotal())

	// Numbers of people
	fs += fmt.Sprintf("%-25v ...%v\n", "numbers of people:", sb.numbersOfPeople)

	// Total per person
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "total per person:", sb.calculateTotalPerPerson())

	// Tip per person
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "tip per person:", sb.calculateTipPerPerson())

	// Share per person
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "share per person:", sb.calculateShare())

	return fs
}

// Ask for numbers of people, total amount and tip
func promptOptionsShareBill(sb *shareBill) {
	var totalAmount float64
	fmt.Println("What is the total amount?")
	fmt.Scanln(&totalAmount)

	sb.totalAmount = totalAmount

	var tip float64
	fmt.Println("What is the tip amount?")
	fmt.Scanln(&tip)

	sb.tip = tip

	var numbersOfPeople int
	fmt.Println("How many people?")
	fmt.Scanln(&numbersOfPeople)

	sb.numbersOfPeople = numbersOfPeople
}

// Functions to calculate the share bill
func (sb *shareBill) calculateShare() float64 {
	return (sb.totalAmount + sb.tip) / float64(sb.numbersOfPeople)
}

func (sb *shareBill) calculateTotal() float64 {
	return sb.totalAmount + sb.tip
}

func (sb *shareBill) calculateTipPercentage() float64 {
	return (sb.tip / sb.totalAmount) * 100
}

func (sb *shareBill) calculateTipPerPerson() float64 {
	return sb.tip / float64(sb.numbersOfPeople)
}

func (sb *shareBill) calculateTotalPerPerson() float64 {
	return sb.calculateTotal() / float64(sb.numbersOfPeople)
}
