package statement

import "fmt"

func thirdPartyConnectDB() {
	// panicは書くべきではない
	panic("Unable to connect database!")
}

func save() {
	defer func() {
		// panicをcatchしてexitさせない
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func Panic() {
	save()
}