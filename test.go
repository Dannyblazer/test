package main

import "fmt"

/*
Question 2: Applying a discount to selected items (~12 minutes)

You have the types below. Implement CalculateDiscount.

Rules:
  1. Only matching lines receive the discount.
  2. For "category", match CategoryID.
  3. For "products", match ProductID.
  4. A percentage discount uses basis points:
       100  = 1%
       2500 = 25%
  5. A fixed discount is in minor units.
  6. The discount cannot exceed the value of the eligible items.
  7. Never use floating point.

Example:
    Item A: 10,000 x 2, category=food
    Item B:  5,000 x 1, category=electronics

    Discount:
      20%
      category=food

    Discount = 4,000
*/

type LineItem struct {
	ProductID  string
	CategoryID string
	UnitPrice  int64
	Quantity   int64
}

type Discount struct {
	Kind      string // "percentage" or "fixed"
	Target    string // "category" or "products"
	TargetIDs []string
	Value     int64 // basis points for percentage, minor units for fixed
}

func CalculateDiscount(
	lines []LineItem,
	discount Discount,
) (int64, error) {

	// Stream items in line
	var final_discount int64
	var itemID string
	if discount.Kind != "fixed" && discount.Value > 10000 {
		return 0, fmt.Errorf("invalid percentage discount specified")
	}
	for _, item := range lines {
		// switch for discount Target types upcoming
		switch discount.Target {
		case "category":
			itemID = item.CategoryID
		case "products":
			itemID = item.ProductID
		default:
			return 0, fmt.Errorf("invalid discount target")
		}

		switch discount.Kind {
		case "percentage":
			for _, targetID := range discount.TargetIDs {
				if itemID == targetID {
					// matched, apply percentage discount
					final_discount += (item.Quantity * item.UnitPrice * discount.Value) / 10000
				}
			}
		case "fixed":
			for _, targetID := range discount.TargetIDs {
				if itemID == targetID {
					// matched, apply percentage discount
					discounted := item.Quantity * discount.Value
					if discounted > item.Quantity*item.UnitPrice {
						final_discount += item.Quantity * item.UnitPrice
					} else {
						final_discount += discounted
					}
				}
			}
		default:
			return 0, fmt.Errorf("invalid discount kind")
		}
	}
	return final_discount, nil
}

func main() {
	fmt.Println("ready")

}
