package utils

func Contains(slice []string, value string) int {
	for idx, ele := range slice {
		if ele==value {
			return idx
		}
	}
	return -1;
}
