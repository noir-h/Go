package statement

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "mac"
}

func Switch() {
	// os := "Mac"
	// os := getOsName()

	// 省略
	switch os := getOsName(); os {
	case "Mac" :{
		fmt.Println("Mac!")
	}
	case "Windows" :{
		fmt.Println("Winsows!")
	}
	default:{
		fmt.Println("Default!")
	}
	}
	// Can not declaration
	// fmt.Println(os)

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning!")
	case t.Hour() < 17:
		fmt.Println("Afternoon!")
	}
}