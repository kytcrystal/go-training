package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unit_map := make(map[string]int)
	unit_map["quarter_of_a_dozen"] = 3
	unit_map["half_of_a_dozen"] = 6
	unit_map["dozen"] = 12
	unit_map["small_gross"]	= 120
	unit_map["gross"] = 144
	unit_map["great_gross"]	= 1728
	return unit_map
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if !exists {
		return exists
	}
	bill[item] += value
	return exists
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// Return false if the given item is not in the bill
	item_value, item_exists := bill[item]
	if !item_exists{
		return item_exists
	}

	// Return false if the given unit is not in the units map.
	unit_value, unit_exists := units[unit]
	if !unit_exists{
		return unit_exists
	}

	// Return false if the new quantity would be less than 0.
	item_value -= unit_value
	if item_value < 0 {
		return false
	}

	if item_value == 0 {
		delete(bill, item)
	} else {
		bill[item] = item_value
	}
	
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	// Return 0 and false if the item is not in the bill.
	value, exists := bill[item]
	if !exists {
		return 0, exists
	}
	// Otherwise, return the quantity of the item in the bill and true.
	return value, exists
}
