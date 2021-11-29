package nsq

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
	"testing"
)

func TestHash(t *testing.T) {

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	h := md5.New()

	// 对hostName做一个hash
	io.WriteString(h, hostname)
	//crc32.ChecksumIEEE 是核心
	defaultID := int64(crc32.ChecksumIEEE(h.Sum(nil)) % 1024)
	fmt.Println(defaultID)
}

const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

func TestConstant(t *testing.T) {

	fmt.Println(Ldate)         //   1
	fmt.Println(Ltime)         //   2
	fmt.Println(Lmicroseconds) //   4
	fmt.Println(Llongfile)     //   8
	fmt.Println(Lshortfile)    //   16
	fmt.Println(LUTC)          //   32
	fmt.Println(Lmsgprefix)    //   64
	fmt.Println(LstdFlags)     //   3

}
