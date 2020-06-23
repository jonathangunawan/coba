package user

import "fmt"

//Error List for User
var (
	UserNotFound = fmt.Errorf("User Not Found")
)

const (
	userLog = "[User]"
)
