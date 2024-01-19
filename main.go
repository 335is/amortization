package main

import (
	"flag"
	"fmt"
	"os"
)

var loanFlag = "loan"
var rateFlag = "rate"
var termFlag = "term"
var extraFlag = "extra"

func main() {
	loan := flag.Float64(loanFlag, 30000.00, "Loan amount in dollars")
	rate := flag.Float64(rateFlag, 5.5, "Interest rate in percentage")
	term := flag.Int(termFlag, 24, "Loan duration in months")
	extra := flag.Float64(extraFlag, 0.0, "Extra monthly principal in dollars")
	flag.Parse()

	printInput(*loan, *rate, *term, *extra)

	r, err := Calculate(*loan, *rate, *term, *extra)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	printSummary(r)

	printSchedule(r)
}

func printInput(loanAmount float64, interestRate float64, termMonths int, extra float64) {
	fmt.Printf("Loan amount:             %s\n", toCurrency(loanAmount))
	fmt.Printf("Interest rate:           %.2f%%\n", interestRate)
	fmt.Printf("Loan term in months:     %d\n", termMonths)
	fmt.Printf("Extra monthly principal: %s\n", toCurrency(extra))
}

func printSummary(r *result) {
	// summary info
	fmt.Printf("Total Interest:          %s\n", toCurrency(r.totalInterest))
	fmt.Printf("Total Paid:              %s\n", toCurrency(r.totalPaid))
	fmt.Println()
}

func printSchedule(r *result) {
	// heading and rows of data
	fmt.Printf("%10s %20s %20s %20s %20s\n", "Row", "Balance", "Payment", "Interest", "Principal")
	for row, payment := range r.payments {
		fmt.Printf("%10d %20s %20s %20s %20s\n", row+1, toCurrency(payment.balance), toCurrency(payment.payment), toCurrency(payment.interest), toCurrency(payment.Principal()))
	}
}

func toCurrency(n float64) string {
	return fmt.Sprintf("$%.2f", n)
}
