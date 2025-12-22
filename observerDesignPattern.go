// Exercise: Stock Price Monitoring System
// Problem Statement:
// You're building a real-time stock price monitoring system where multiple investors and services need to be notified whenever a stock price changes. Different observers react differently to price changes.
// Requirements:

// Implement an Observer pattern where a Stock (subject) notifies multiple observers when its price changes
// Create at least 3 different types of observers:

// EmailAlert: Sends email notification when price changes by more than 5%
// MobileApp: Pushes notification to mobile app for any price change
// Logger: Logs all price changes to console with timestamp

// Observers should be able to:

// Subscribe to a stock
// Unsubscribe from a stock
// Receive updates with old price and new price

// The Stock should:

// Allow setting a new price
// Automatically notify all registered observers when price changes
// Keep track of its current price and symbol (e.g., "AAPL", "GOOGL")

// Bonus Challenges:

// Implement a PriceThresholdObserver that only triggers when price crosses a specific threshold
// Add support for multiple stocks, where observers can subscribe to different stocks
// Implement a TradingBot observer that automatically "buys" when price drops 10% and "sells" when it rises 10%
// Make the notification system concurrent-safe using goroutines and channels

// Expected Output Example:
// [LOGGER] AAPL price changed: $150.00 -> $157.50 (+5.00%)
// [EMAIL] Alert! AAPL price jumped by 5.00%
// [MOBILE] Push notification: AAPL is now $157.50

// [LOGGER] AAPL price changed: $157.50 -> $160.00 (+1.59%)
// [MOBILE] Push notification: AAPL is now $160.00

// [MOBILE] Unsubscribed from AAPL updates

// [LOGGER] AAPL price changed: $160.00 -> $155.00 (-3.12%)
// Hints:

// Define an Observer interface with an Update(stock string, oldPrice, newPrice float64) method
// The Stock struct should maintain a slice of observers
// Think about how to handle observer registration/removal
// Consider using sync.Mutex if you want to make it thread-safe

// Interview Follow-up Questions to Consider:

// What happens if an observer's Update() method is slow? How would you handle that?
// How is this different from a simple callback mechanism?
// What are the advantages of Observer pattern over tight coupling?
// How would you prevent memory leaks if observers aren't properly unsubscribed?
// Could you implement this using Go channels instead? What would be the trade-offs?

// Try implementing this yourself first, then I can review your solution or show you an example implementation. Good luck! 📈
