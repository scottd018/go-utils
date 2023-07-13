package list

// Has determines if a list contains an integer.
func (list integers) Has(integer int) bool {
	if len(list) == 0 {
		return false
	}

	for item := range list {
		if integer == list[item] {
			return true
		}
	}

	return false
}

// Unique gets the unique integers in a list of integers.
func (list integers) Unique() []int {
	if len(list) == 0 {
		return list
	}

	uniqueMap := make(map[int]bool)

	uniqueList := []int{}

	for item := range list {
		if !uniqueMap[list[item]] {
			uniqueList = append(uniqueList, list[item])

			uniqueMap[list[item]] = true
		}
	}

	return uniqueList
}
