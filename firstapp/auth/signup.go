package auth

type User struct {
    Name     string
    Password string
}

func Signup(name, password string) bool {

    if name == "" || password == "" {
        return false
    }

    user := User{
        Name:     name,
        Password: password,
    }

    _ = user // just to avoid unused variable while learning

    return true
}