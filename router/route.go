package router

import "niu/util/array"

var validMethods [5]string = [5]string {"GET", "POST", "HEAD", "PUT", "DELETE"}

// is a valid request method
func isValidMethod(method string) bool {
	return array.InStringArray(method, validMethods[:])

}
