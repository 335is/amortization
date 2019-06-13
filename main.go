package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	loan := flag.Float64("loan", 30000.00, "Loan amount in dollars")
	rate := flag.Float64("rate", 5.5, "Interest rate in percentage")
	term := flag.Int("term", 24, "Loan duration in months")
	extra := flag.Float64("extra", 0.0, "Extra monthly principal in dollars")
	flag.Parse()

	r, err := calculate(*loan, *rate, *term, *extra)
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
