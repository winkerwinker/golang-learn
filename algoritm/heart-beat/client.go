package heart_beat

import (
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"strconv"
	"time"
)

// 发送数据
func client(conn *net.TCPConn) {

	for i := 0; i < 10; i++ {
		words := strconv.Itoa(i) + " Hello I'm MyHeartbeat Client."
		msg, err := conn.Write([]byte(words))
		if err != nil {
			Log(conn.RemoteAddr().String(), "Fatal error: ", err)
			os.Exit(1)
		}
		Log("服务端接收了", msg)
		time.Sleep(2 * time.Second)
	}
	for i := 0; i < 2; i++ {
		time.Sleep(12 * time.Second)
	}
	for i := 0; i < 10; i++ {
		words := strconv.Itoa(i) + " Hi I'm MyHeartbeat Client."
		msg, err := conn.Write([]byte(words))
		if err != nil {
			Log(conn.RemoteAddr().String(), "Fatal error: ", err)
			os.Exit(1)
		}
		Log("服务端接收了", msg)
		time.Sleep(2 * time.Second)
	}
}

func Log(format string, msg ...interface{}) {
	logrus.Info(" format ", msg)
}
