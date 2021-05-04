package utils

func Contains(slice []string, value string) int {
	for idx, ele := range slice {
		if ele == value {
			return idx
		}
	}
	return -1
}

func Replace(slice []string, from string, to string) []string {
	sliceCpy := make([]string, len(slice))
	copy(sliceCpy, slice)
	for idx, val := range sliceCpy {
		if val == from {
			sliceCpy[idx] = to
		}
	}
	return sliceCpy
}
