package statement

import (
	"fmt"
	"io"
	"log"
	"os"
)

func loggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// MultiWriter は様々な出力先を提供する
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

func Log() {
	loggingSettings("test.log")
	_, err := os.Open("fdsfdsf")
	if err != nil {
		log.Fatalln("Exit", err)
	}
	log.Println("logging")
	log.Printf("%T %v", "test", "test")

	// exitする
	log.Fatalf("%T %v", "test", "test")
	log.Fatalln("error!")

	fmt.Println("ok!")
}