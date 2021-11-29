package heart_beat

import (
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"time"
)

// 用于接受数据，多少秒一个客户端没有数据就关闭链接
func server() {

	server := ":7373"
	netListen, err := net.Listen("tcp", server)
	if err != nil {
		logrus.Info("connect error: ", err)
		os.Exit(1)
	}
	logrus.Info("Waiting for Client ...")

	for {
		// 可以不停的建立链接吗
		conn, err := netListen.Accept()
		if err != nil {
			// 简历连接失败
			logrus.Info(conn.RemoteAddr().String(), "Fatal error: ", err)
			continue
		}

		// 设置10s 超时
		conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))

		logrus.Info(conn.RemoteAddr().String(), "connect success!")

	}

}

func handleCoon(conn net.Conn) {
	// todo
	buffer := make([]byte, 1024)
	defer conn.Close()

	for {

		read, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " Fatal error: ", err)
			return
		}
		Data := buffer[:read]
		message := make(chan byte)
		//心跳计时
		go HeartBeating(conn, message, 4)
		//检测每次是否有数据传入
		go GravelChannel(Data, message)

		//		Log(time.Now().Format("2006-01-02 15:04:05.0000000"), conn.RemoteAddr().String(), string(buffer[:n]))
	}

}

func GravelChannel(bytes []byte, mess chan byte) {
	for _, v := range bytes {
		mess <- v
	}
	close(mess)
}

// 心跳
func HeartBeating(conn net.Conn, bytes chan byte, timeout int) {
	select {
	case fk := <-bytes:
		Log(conn.RemoteAddr().String(), "心跳:第", string(fk), "times")
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break

	case <-time.After(5 * time.Second):
		Log("conn dead now")
		conn.Close()
	}
}
