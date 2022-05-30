package statement

import (
	"fmt"
	"log"
	"os"
)

func Error() {
	file, err := os.Open("./main.go")
	if err != nil {
		log.Fatalln("Error!")
	} 
	defer file.Close()
	data := make([]byte, 100)
	// shor declarationは何かしら一つinitialできればerrorにならない
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error!")
	}
	fmt.Println(count, string(data))

	// 返り値がerror一つだけとかはinitialできない, overrideする必要がある 
	// err = os.Chdir("fdsf")
	// err一つの返り値の時、よくこういう書き方をする 
	if err = os.Chdir("fsddfds"); err != nil {
		log.Fatalln("Error!")
	}
}