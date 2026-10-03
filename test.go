package main

import (
	"fmt"
)

// /*
// Question 2: Applying a discount to selected items (~12 minutes)

// You have the types below. Implement CalculateDiscount.

// Rules:
//   1. Only matching lines receive the discount.
//   2. For "category", match CategoryID.
//   3. For "products", match ProductID.
//   4. A percentage discount uses basis points:
//        100  = 1%
//        2500 = 25%
//   5. A fixed discount is in minor units.
//   6. The discount cannot exceed the value of the eligible items.
//   7. Never use floating point.

// Example:
//     Item A: 10,000 x 2, category=food
//     Item B:  5,000 x 1, category=electronics

//     Discount:
//       20%
//       category=food

//     Discount = 4,000
// */

// type LineItem struct {
// 	ProductID  string
// 	CategoryID string
// 	UnitPrice  int64
// 	Quantity   int64
// }

// type Discount struct {
// 	Kind      string // "percentage" or "fixed"
// 	Target    string // "category" or "products"
// 	TargetIDs []string
// 	Value     int64 // basis points for percentage, minor units for fixed
// }

// func CalculateDiscount(
// 	lines []LineItem,
// 	discount Discount,
// ) (int64, error) {

// 	// Stream items in line
// 	var final_discount int64
// 	var itemID string
// 	if discount.Kind != "fixed" && discount.Value > 10000 {
// 		return 0, fmt.Errorf("invalid percentage discount specified")
// 	}
// 	for _, item := range lines {
// 		// switch for discount Target types upcoming
// 		switch discount.Target {
// 		case "category":
// 			itemID = item.CategoryID
// 		case "products":
// 			itemID = item.ProductID
// 		default:
// 			return 0, fmt.Errorf("invalid discount target")
// 		}

// 		switch discount.Kind {
// 		case "percentage":
// 			for _, targetID := range discount.TargetIDs {
// 				if itemID == targetID {
// 					// matched, apply percentage discount
// 					final_discount += (item.Quantity * item.UnitPrice * discount.Value) / 10000
// 				}
// 			}
// 		case "fixed":
// 			for _, targetID := range discount.TargetIDs {
// 				if itemID == targetID {
// 					// matched, apply percentage discount
// 					discounted := item.Quantity * discount.Value
// 					if discounted > item.Quantity*item.UnitPrice {
// 						final_discount += item.Quantity * item.UnitPrice
// 					} else {
// 						final_discount += discounted
// 					}
// 				}
// 			}
// 		default:
// 			return 0, fmt.Errorf("invalid discount kind")
// 		}
// 	}
// 	return final_discount, nil
// }

// Exercise 2

// Question 3: Deduplicating and Settling a Batch of Transactions (~12 minutes)

// You have the types below. Implement SettleBatch.

// Rules:

// 1. Transactions can arrive with exact duplicate IdempotencyKey values, because a client may retry a request that already succeeded.
// 2. When two or more transactions share an IdempotencyKey, only the first one encountered counts. Every later duplicate with that same key is ignored entirely, it does not affect the balance at all.
// 3. ach transaction's Type is either "credit" or "debit".
// 4. The net balance is the sum of credits minus the sum of debits, counting only the de-duplicated transactions.
// 5. If the resulting net balance would be negative, return an error instead of a negative number.
// 6. Amounts are in minor units (integers). Never use floating point.

// type Transaction struct {
// 	IdempotencyKey string
// 	Type           string // "credit" or "debit"
// 	Amount         int64
// }

// func SettleBatch(transactions []Transaction) (int64, error) {
// 	// implement
// 	var net_balance int64
// 	processed_keys := make(map[string]string)

// 	for _, transaction := range transactions {
// 		fmt.Println("Processing ID: ", transaction.IdempotencyKey)

// 		if _, seen := processed_keys[transaction.IdempotencyKey]; !seen {
// 			processed_keys[transaction.IdempotencyKey] = transaction.IdempotencyKey
// 			// process the transaction
// 			switch transaction.Type {
// 			case "credit":
// 				net_balance += transaction.Amount
// 			case "debit":
// 				net_balance -= transaction.Amount
// 			default:
// 				fmt.Println("invalid transaction type for trnxID: ", transaction.IdempotencyKey)
// 			}
// 		}

// 	}
// 	if net_balance < 0 {
// 		return 0, fmt.Errorf("negative balance err")
// 	}
// 	return net_balance, nil
// }

// 3
// Rules:

// Only "debit" transactions count toward the limit. "credit" transactions are ignored entirely, not even counted as noise.
// The input is not guaranteed to be sorted by timestamp.
// A violation occurs if any debit transaction t has at least maxCount debit transactions (including itself) whose timestamps fall within the inclusive range [t - windowSeconds, t].
// Return true and the timestamp of the transaction that triggered the violation, as soon as you find one. If no violation exists anywhere in the batch, return false and 0.
// Never use floating point. All timestamps and windows are integers already.

// Example:

// Debit timestamps: 100, 105, 110, 300
// maxCount = 3, windowSeconds = 15

// At t=110: the window is [95, 110]. Debits at 100, 105, and 110 all fall inside it — that's 3, meeting maxCount.
// Violation triggers at t=110.

// Return: (true, 110)

type Transaction struct {
	Type      string // "credit" or "debit"
	Timestamp int64  // unix seconds
}

// func ExceedsVelocityLimit(transactions []Transaction, maxCount int, windowSeconds int64) (bool, int64) {
// 	// implement
// 	// First try to sort the []Transaction in ascending order
// 	debitTimes := make([]int64, 0, len(transactions))
// 	for _, t := range transactions {
// 		if t.Type == "debit" {
// 			debitTimes = append(debitTimes, t.Timestamp)
// 		}
// 	}

// 	sort.Slice(debitTimes, func(i, j int) bool { return debitTimes[i] < debitTimes[j] })
// 	left := 0
// 	// loop through the transactions
// 	for right, t := range debitTimes {
// 		windowsEdge := debitTimes[right] - windowSeconds
// 		for debitTimes[left] < windowsEdge {
// 			//fmt.Println("Left forward")
// 			left++
// 		}
// 		//fmt.Printf("left: %v and right: %v\n", left, right)
// 		if right-left+1 >= maxCount {
// 			return true, t
// 		}

// 	}
// 	return false, 0

// }
func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing work!")
		}
	}
}
func main() {
	// 	transactions := []Transaction{
	// 		//{Type: "debit", Timestamp: 90},
	// 		{Type: "debit", Timestamp: 100},
	// 		{Type: "debit", Timestamp: 105},
	// 		{Type: "debit", Timestamp: 110},
	// 		// {Type: "credit", Timestamp: 115},
	// 		// {Type: "debit", Timestamp: 120},
	// 		// {Type: "credit", Timestamp: 121},
	// 		// {Type: "debit", Timestamp: 122},
	// 		// {Type: "debit", Timestamp: 123},
	// 		{Type: "debit", Timestamp: 124},
	// 	}
	// 	status, value := ExceedsVelocityLimit(transactions, 3, 13)
	// 	fmt.Printf("Status: %v and Value: %v\n", status, value)

	chars := []string{"a", "b", "c"}
	// myChannel := make(chan string)
	// nextChannel := make(chan string)
	channelChars := make(chan string, 3)

	// go func() {
	// 	nextChannel <- "next please"
	// }()
	// go func() {
	// 	myChannel <- "hello"
	// }()

	for _, s := range chars {
		//fmt.Println("sent: ", s)

		channelChars <- s
	}

	// select {
	// case myChannelMsg := <-myChannel:
	// 	fmt.Println(myChannelMsg)
	// case nextChannelMsg := <-nextChannel:
	// 	fmt.Println(nextChannelMsg)
	// }

	// done := make(chan bool)

	// go doWork(done)
	// //time.Sleep(1 * time.Second)

	// close(done)
	close(channelChars)

	for data := range channelChars {
		fmt.Println("data: ", data)
	}

	fmt.Println("Ready")

}
