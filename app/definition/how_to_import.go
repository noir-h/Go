package definition

import (
	"fmt"
	"os/user"
	"time"
)

func How_to_import(){
	fmt.Println("Hello", time.Now())
	fmt.Println(user.Current())
}