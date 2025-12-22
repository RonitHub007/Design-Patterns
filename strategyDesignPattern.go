/*
Exercise: Payment Processing System
Problem Statement:
You're building an e-commerce checkout system that needs to support multiple payment methods
(Credit Card, PayPal, Cryptocurrency). Each payment method has different processing logic and validation rules.
Requirements:

Implement a Strategy pattern where each payment method is a separate strategy
Create a PaymentContext that can switch between different payment strategies
Each payment strategy should:

Validate the payment details
Process the payment
Return a transaction ID and any errors

Implement at least 3 different payment strategies:

CreditCard: Requires card number (16 digits), CVV (3 digits), expiry date
PayPal: Requires email address validation
Crypto: Requires wallet address (starts with "0x" and is 42 chars long)

Bonus Challenges:

Add a method to calculate processing fees (different for each strategy)
Implement a Refund() method with strategy-specific logic
Add logging/tracking for each payment attempt

Expected Output Example:
Processing Credit Card payment of $100.50...
✓ Payment successful! Transaction ID: CC-20241217-1234

Processing PayPal payment of $50.00...
✓ Payment successful! Transaction ID: PP-20241217-5678
Hints:

Define a PaymentStrategy interface with methods like Pay(amount float64) (string, error)
Think about what data each strategy needs (hint: struct fields)
Consider using method receivers for strategy implementations

Interview Follow-up Questions to Consider:

Why use Strategy pattern instead of a large if-else or switch statement?
How would you add a new payment method without modifying existing code?
What are the trade-offs of this pattern?
*/
package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

/*
Strategy Interface
*/
type PaymentStrategy interface {
	Pay(amount float64) (string, error)
	Refund(amount float64) (string, error)
	CalculateProcessingFee(amount float64) (float64, error)
}

/*
Payment Context (Strategy Switcher)
*/
type PaymentContext struct {
	strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) Pay(amount float64) {
	fmt.Printf("\nProcessing %T payment of $%.2f...\n", p.strategy, amount)

	tx, err := p.strategy.Pay(amount)
	if err != nil {
		fmt.Println("❌ Payment failed:", err)
		return
	}

	fee, _ := p.strategy.CalculateProcessingFee(amount)
	fmt.Printf("✓ Payment successful! %s\n", tx)
	fmt.Printf("Processing Fee: $%.2f\n", fee)
}

/*
Concrete Strategies
*/
type CreditCard struct {
	CardNumber string
	CVV        string
	ExpiryDate string
}

type Paypal struct {
	Email string
}

type Crypto struct {
	WalletAddress string
}

/*
Utility
*/
func transactionIdGenerator(paymentType string) string {
	return fmt.Sprintf(
		"%s-%s-%d",
		paymentType,
		time.Now().Format("20060102"),
		time.Now().UnixNano(),
	)
}

/*
Credit Card Implementation
*/
func (c *CreditCard) validate() error {
	if len(c.CardNumber) != 16 {
		return errors.New("invalid card number")
	}
	if len(c.CVV) != 3 {
		return errors.New("invalid CVV")
	}
	if c.ExpiryDate == "" {
		return errors.New("expiry date missing")
	}
	return nil
}

func (c *CreditCard) Pay(amount float64) (string, error) {
	if err := c.validate(); err != nil {
		return "", err
	}
	return transactionIdGenerator("CC"), nil
}

func (c *CreditCard) Refund(amount float64) (string, error) {
	return transactionIdGenerator("CC-REFUND"), nil
}

func (c *CreditCard) CalculateProcessingFee(amount float64) (float64, error) {
	return amount * 0.01, nil
}

/*
Paypal Implementation
*/
func (p *Paypal) validate() error {
	if !strings.Contains(p.Email, "@") {
		return errors.New("invalid PayPal email")
	}
	return nil
}

func (p *Paypal) Pay(amount float64) (string, error) {
	if err := p.validate(); err != nil {
		return "", err
	}
	return transactionIdGenerator("PP"), nil
}

func (p *Paypal) Refund(amount float64) (string, error) {
	return transactionIdGenerator("PP-REFUND"), nil
}

func (p *Paypal) CalculateProcessingFee(amount float64) (float64, error) {
	return amount * 0.02, nil
}

/*
Crypto Implementation
*/
func (c *Crypto) validate() error {
	if !strings.HasPrefix(c.WalletAddress, "0x") || len(c.WalletAddress) != 42 {
		return errors.New("invalid crypto wallet address")
	}
	return nil
}

func (c *Crypto) Pay(amount float64) (string, error) {
	if err := c.validate(); err != nil {
		return "", err
	}
	return transactionIdGenerator("CRYPTO"), nil
}

func (c *Crypto) Refund(amount float64) (string, error) {
	return transactionIdGenerator("CRYPTO-REFUND"), nil
}

func (c *Crypto) CalculateProcessingFee(amount float64) (float64, error) {
	return amount * 0.03, nil
}

/*
Main Processor
*/
func PaymentProcessor() {
	context := &PaymentContext{}

	creditCard := &CreditCard{
		CardNumber: "1234567890123456",
		CVV:        "123",
		ExpiryDate: "12/26",
	}

	paypal := &Paypal{
		Email: "ronitroy@example.com",
	}

	crypto := &Crypto{
		WalletAddress: "0x1234567890123456789012345678901234567890",
	}

	context.SetStrategy(creditCard)
	context.Pay(100.50)

	context.SetStrategy(paypal)
	context.Pay(50.00)

	context.SetStrategy(crypto)
	context.Pay(75.25)
}
