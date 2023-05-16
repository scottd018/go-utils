package list

// Has determines if a list contains a string.
func (list strings) Has(str string) bool {
	if len(list) == 0 {
		return false
	}

	for item := range list {
		if str == list[item] {
			return true
		}
	}

	return false
}

// Unique gets the unique strings in a list of strings.
func (list strings) Unique() []string {
	if len(list) == 0 {
		return list
	}

	uniqueMap := make(map[string]bool)

	uniqueList := []string{}

	for item := range list {
		if !uniqueMap[list[item]] {
			uniqueList = append(uniqueList, list[item])

			uniqueMap[list[item]] = true
		}
	}

	return uniqueList
}
