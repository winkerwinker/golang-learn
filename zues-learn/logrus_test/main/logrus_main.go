package main

import (
	logger "awesomeProject/zues-learn/logger"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

func main() {
	//log1()
	//log2()
	logFormat()
}




type MyFormatter struct{}

func (s *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	msg := fmt.Sprintf("%s [%s] %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

func logFormat() {

	logrus.SetFormatter(new(MyFormatter))
	logrus.SetOutput(os.Stdout)
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}

func log2() {
	fmt.Printf("\n\n")
	logger.Logger.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}

func log1() {
	fmt.Printf("\n\n")
	//INFO[0000] A walrus appears                              animal=walrus
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}


