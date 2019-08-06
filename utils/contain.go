package utils

func ContainString(target string, list []string) bool{

	for _, elem := range list {
		if target == elem {
			return true
		}
	}
	return false
}
