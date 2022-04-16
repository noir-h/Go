package how_to_import

import (
	"fmt"
	"os/user"
	"time"
)

func Test(){
	fmt.Println("Hello", time.Now())
	fmt.Println(user.Current())
}