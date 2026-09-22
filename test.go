package main

import "fmt"

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

type Transaction struct {
	IdempotencyKey string
	Type           string // "credit" or "debit"
	Amount         int64
}

func SettleBatch(transactions []Transaction) (int64, error) {
	// implement
	var net_balance int64
	processed_keys := make(map[string]string)

	for _, transaction := range transactions {
		fmt.Println("Processing ID: ", transaction.IdempotencyKey)
		// if _, seen := processed_keys[transaction.IdempotencyKey]; seen {
		// 	// Skip duplicate transaction and exit processedKeys looping
		// 	// processed_keys[transaction.IdempotencyKey]
		// 	fmt.Println("Duplicate transaction skipped")
		// 	break
		// }
		if _, seen := processed_keys[transaction.IdempotencyKey]; !seen {
			processed_keys[transaction.IdempotencyKey] = transaction.IdempotencyKey
			// process the transaction
			switch transaction.Type {
			case "credit":
				net_balance += transaction.Amount
			case "debit":
				net_balance -= transaction.Amount
			default:
				fmt.Println("invalid transaction type for trnxID: ", transaction.IdempotencyKey)
			}
		}

	}
	if net_balance < 0 {
		return 0, fmt.Errorf("negative balance err")
	}
	return net_balance, nil
}

func main() {
	transactions := []Transaction{
		{IdempotencyKey: "abcd", Type: "credit", Amount: 10000},
		{IdempotencyKey: "abcd", Type: "debit", Amount: 10000},
		{IdempotencyKey: "abcdx", Type: "credit", Amount: 50000},
	}
	fmt.Println("ready")
	balance, err := SettleBatch(transactions)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(balance)
}
