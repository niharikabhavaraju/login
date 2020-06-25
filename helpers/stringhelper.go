package helpers

//IsEmpty checks if string is empty
func IsEmpty(data string) bool {
	if len(data) <= 0 {
		return true
	} else {
		return false
	}

}
