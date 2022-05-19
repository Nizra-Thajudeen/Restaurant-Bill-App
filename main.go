package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	myBill := createBill()
	getOption(myBill)

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Enter bill name :", reader)

	b := newBill(name)

	return b
}

func getInput(question string, r *bufio.Reader) (string, error) {
	fmt.Println(question)
	inputVal, err := r.ReadString('\n')

	return strings.TrimSpace(inputVal), err
}

func getOption(b bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Select option (a - add an item, s - save bill, t - add tip)", reader)

	switch option {
	case "a":
		name, _ := getInput("Item name : ", reader)
		price, _ := getInput("Item price : ", reader)

		cPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("price should be a number")
			getOption(b)
		}

		b.addItem(name, cPrice)
		fmt.Println("Item added")
		getOption(b)

	case "t":
		tip, _ := getInput("Enter tip amount in $ : ", reader)

		cTip, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("tip should be a number")
			getOption(b)
		}

		b.updateTips(cTip)

		fmt.Println("tip added")
		getOption(b)

	case "s":
		b.saveBill()
		fmt.Println("you saved the bill ", b.name)
		fmt.Println(b)
	default:
		fmt.Println("choose a valid option")
		getOption(b)
	}
}
