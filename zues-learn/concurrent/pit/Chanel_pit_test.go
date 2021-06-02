package chanel

import (
	"testing"
)

// golang 的坑
func TestLogrue(t *testing.T) {

}


package main

import (
log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}