package repos

//UserIsValid checks if user is valid
func UserIsValid(uName, pwd string) bool {
	_uName, _pwd, _isValid := "niharika", "1234", false

	if uName == _uName && pwd == _pwd {
		_isValid = true
	} else {
		_isValid = false
	}
	return _isValid
}
