package hello

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var myLogger *log.Logger

func TestLogger(t *testing.T) {
	fpLog, err := os.OpenFile("C:/goClass/iotest/logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}
	fmt.Println(err)
	defer fpLog.Close()

	myLogger = log.New(fpLog, "Info : ", log.Ldate|log.Ltime|log.Lshortfile)
	myLogger.Println("test")

	myLogger.Println("End")
}
