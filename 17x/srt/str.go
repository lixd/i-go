package srt

func IsValidUsername(username string) bool {
	if username == "" {
		return false
	}
	return true
}
func IsValidPassword(password string) bool {
	if password == "" {
		return false
	}
	return true
}

func IsValidUsernameOrPassword(str string) bool {
	if str == "" {
		return false
	}
	return true
}
func IsValidIp(ip string) bool {
	if ip == "" {
		return false
	}
	return true
}
func CheckIpIfValid(ip string) bool {
	if len(ip) == 0 {
		return false
	}
	return true
}

/*func Login(email, password string) {
	if IsUserExists(email, password) {
		GetUserByEmail(email)
	}
}
func IsUserExists(email, password string) bool {
	if isValidEmail(email) && isValidPassword(password) {
		return true
	}
	// query db
	return false
}
func GetUserByEmail(email string) {
	if isValidEmail(email) {
		//	query db
	}
}*/
