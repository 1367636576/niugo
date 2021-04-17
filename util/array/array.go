package array

//a num is in inter array
func InIntArray(val int,arr []int) bool {
	for _,item := range arr{
		if item == val {
			return true
		}
	}
	return false
}

// a string is in string array
func InStringArray(val string,arr []string) bool {
	for _,item := range arr{
		if item == val {
			return true
		}
	}
	return false
}
