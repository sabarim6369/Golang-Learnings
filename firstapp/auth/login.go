package auth
func Login(name string, password string) bool {
	if(name=="admin" && password=="admin"){
		return true
	}
	return false


}