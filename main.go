package main

import "fmt"

func main() {
	fmt.Println("1. Normal bill")
	fmt.Println("2. Share bill")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		mybill := createBill()
		promptOptions(mybill)
		fmt.Println(mybill.format())
	case 2:
		shareBill := createShareBill()
		promptOptionsShareBill(&shareBill)
		fmt.Println(shareBill.formatShareBill())
		saveShareBill(&shareBill)
	default:
		fmt.Println("Invalid option, please try again")
	}
}
