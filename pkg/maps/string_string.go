package maps

// HasKey determines if a map has a key.
func (m strings) HasKey(key string) bool {
	if len(m) == 0 {
		return false
	}

	for k := range m {
		if k == key {
			return true
		}
	}

	return false
}

// HasValue determines if a map has a value.
func (m strings) HasValue(value string) bool {
	if len(m) == 0 {
		return false
	}

	for k := range m {
		if m[k] == value {
			return true
		}
	}

	return false
}
