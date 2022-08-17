package utils

func ArrayContainsUint(arr []uint, value uint) bool {
	for _, element := range arr {
		if element == value {
			return true
		}
	}

	return false
}
