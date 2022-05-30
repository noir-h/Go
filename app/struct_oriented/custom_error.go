package structoriented

import "fmt"

type UserNotFound struct {
	Username string 
}

// printlnした時に出力内容を上書きしたい時に使う
// お作法で引数はポインタ型
func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found %v", e.Username)
}

func myFunc() error {
	// something wrong
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "sora"}
}
func CustomError() {
	e1 := &UserNotFound{"sora"}
	fmt.Println(e1)
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}