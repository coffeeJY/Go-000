package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

var nams = []string{
	"Jack",
	"K",
	"Jon",
	"HelloKiKi",
}

type HelloMsg struct {
	Id  int
	Msg string
}

func main() {
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8889")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	log.Println("dial ok")
	defer func() {
		conn.Close()
		log.Println("close ok")
	}()

	var (
		rd    = rand.New(rand.NewSource(time.Now().UnixNano()))
		index int
	)
	var sendByteNum int
	for i := 0; i < 100; i++ {
		log.Println("Write数据")
		index = rd.Intn(len(nams))
		var msg = HelloMsg{
			i + 1,
			fmt.Sprintf("My name is %s", nams[index]),
		}
		b, err := json.Marshal(msg)
		if err != nil {
			fmt.Printf("json Marshal err : %v \n", err)
		}
		log.Printf("b :%v, len : %v \n", b, len(b))
		n, err := conn.Write(b)
		if err != nil {
			log.Println("write error:", err)
		} else {
			log.Printf("write %v bytes, content is %v\n", n, b)
			sendByteNum += n
		}
		time.Sleep(time.Second * 1)
	}
	log.Printf("sendByteNum : %v \n", sendByteNum)
	time.Sleep(time.Second * 1000)
}
