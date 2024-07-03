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
	panic("Please implement the AddItem() function")
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	panic("Please implement the RemoveItem() function")
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	panic("Please implement the GetItem() function")
}
