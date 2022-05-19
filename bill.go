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

//make new bills
func newBill(name string) bill {
	bills := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return bills
}

//format the bill
func (b *bill) formatBill() string {
	formatted := "Bill Breakdown \n"
	total := 0.0

	for key, value := range b.items {
		formatted += fmt.Sprintf("%-25v ...$%v \n", key+":", value)
		total += value
	}

	//add tip
	formatted += fmt.Sprintf("%-25v ...$%v", "tip: ", b.tip)
	//total
	formatted += fmt.Sprintf("%-25v ...$%0.2f", "total: ", total+b.tip)

	return formatted
}

//update tips
func (b *bill) updateTips(tip float64) {
	b.tip = tip
}

//add items
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

//save bill
func (b *bill) saveBill() {
	data := []byte(b.formatBill())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0466)

	if err != nil {
		panic(err)
	}

	fmt.Println("The bill is saved to file")

}
