package main

import (
	"fmt"
	"math"
)

type result struct {
	payments       []payment
	monthlyPayment float64
	totalInterest  float64
	totalPaid      float64
}

type payment struct {
	balance  float64
	payment  float64
	interest float64
}

func (p *payment) Principal() float64 {
	return p.payment - p.interest
}

func (p *payment) Equity(loanAmount float64) float64 {
	return loanAmount - p.balance
}

func Calculate(loanAmount float64, interestRate float64, termMonths int, extra float64) (*result, error) {
	if loanAmount <= 0.0 {
		return nil, fmt.Errorf("invalid loan amount %f dollars, must be > 0", loanAmount)
	}

	if interestRate < 0.0 {
		return nil, fmt.Errorf("invalid interest rate %f percent, must be >= 0", interestRate)
	}

	if termMonths <= 0.0 {
		return nil, fmt.Errorf("invalid loan term %d months, must be > 0", termMonths)
	}

	if extra < 0.0 {
		return nil, fmt.Errorf("invalid extra principal %f dollars, must be >= 0", extra)
	}

	monthlyInterestRate := (interestRate / 100.0) / 12.0

	r := result{
		payments:       []payment{},
		monthlyPayment: paymentAmount(loanAmount, monthlyInterestRate, float64(termMonths)) + extra,
		totalInterest:  0.0,
		totalPaid:      0.0,
	}

	b := loanAmount
	last := false
	for row := 0; row < termMonths && !last; row++ {
		p := payment{}
		p.interest = b * monthlyInterestRate
		p.payment = r.monthlyPayment
		if r.monthlyPayment > b+p.interest {
			// special case the last payment
			p.payment = b + p.interest
			last = true
		}
		p.balance = b

		b -= (p.payment - p.interest)

		r.payments = append(r.payments, p)
		r.totalInterest += p.interest
		r.totalPaid += p.payment
	}

	return &r, nil
}

func paymentAmount(loanAmount float64, monthlyInterestRate float64, months float64) float64 {
	p := math.Pow(monthlyInterestRate+1, months)
	return loanAmount * (monthlyInterestRate * p) / (p - 1)
}
