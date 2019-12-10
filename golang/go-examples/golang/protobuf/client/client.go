package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
)

func main() {
	strIP := "localhost:6600"
	var conn net.Conn
	var err error

	for conn, err = net.Dial("tcp", strIP); err != nil; conn, err = net.Dial("tcp", strIP) {
		fmt.Println("connect", strIP, "fail")
		time.Sleep(time.Second)
		fmt.Println("reconnect...")
	}
	fmt.Println("connect", strIP, "success")
	defer conn.Close()

	cnt := 0
	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
		cnt++
		stSend := &UserInfo{
			Message: sender.Text(),
			Length:  *proto.Int(len(sender.Text())),
			Cnt:     *proto.Int(cnt),
		}

		pData, err := proto.Marshal(stSend)
		if err != nil {
			panic(err)
		}

		conn.Write(pData)
		if sender.Text() == "stop" {
			return
		}
	}
}
