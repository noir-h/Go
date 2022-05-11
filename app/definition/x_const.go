package definition

import "fmt"

const Pi = 3.14

const(
	username = "test_user"
	password = "test_pass"
)

// case of overflow, stack?がoverflow
// var big int = 9223372036854775807 + 1
// const big = 9223372036854775807 + 1

func Const(){
	fmt.Println(Pi, username, password)
}