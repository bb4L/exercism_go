package gross

var allUnits = map[string]int{
	"quarter_of_a_dozen": 3,
	"half_of_a_dozen":    6,
	"dozen":              12,
	"small_gross":        120,
	"gross":              144,
	"great_gross":        1728,
}

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return allUnits
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	count, ok := units[unit]

	if ok {
		existingCount := bill[item]
		bill[item] = count + existingCount
	}

	return ok
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	val, ok := units[unit]

	if !ok {
		return ok
	}

	count, ok2 := bill[item]

	if !ok2 {
		return ok2
	}

	if val > count {
		return false
	}

	if val == count {
		delete(bill, item)
	} else {
		bill[item] = count - val
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	val, ok := bill[item]
	return val, ok
}
