package main

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
