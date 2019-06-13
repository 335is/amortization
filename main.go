package main

import (
	"fmt"
	"os"
)

func main() {
	loan := 30000.00
	rate := 5.5
	term := 24
	extra := 1000.0

	r, err := calculate(loan, rate, term, extra)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	for row, payment := range r.payments {
		fmt.Printf("%d %s %s %s %s\n", row+1, toCurrency(payment.balance), toCurrency(payment.payment), toCurrency(payment.interest), toCurrency(payment.Principal()))
	}
}

func toCurrency(n float64) string {
	return fmt.Sprintf("$%.2f", n)
}
