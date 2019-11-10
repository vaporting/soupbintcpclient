package main

import (
	"fmt"

	"time"

	"os"

	"os/signal"

	"flag"

	"soupbintcpclient/types"

	"soupbintcpclient/producer"

	"soupbintcpclient/net"
)

func receive(sigC chan os.Signal, stopFlag *bool) {
	for {
		select {
		case <-sigC:
			*stopFlag = true
			break
		}
	}
}

func produceCmds(cmdr *producer.Commander) {
	pkt, _ := types.NewPacket(types.PktTCliHB, []byte(""))
	msg := &types.Message{Msec: 0, Pkt: pkt}
	cmdr.PushMsg(msg)
	pkt, _ = types.NewPacket(types.PktTDbg, []byte("this is debug message"))
	msg = &types.Message{Msec: 1300, Pkt: pkt}
	cmdr.PushMsg(msg)
	pkt, _ = types.NewPacket(types.PktTUData, []byte("this is unsequenced packet 1"))
	msg = &types.Message{Msec: 4600, Pkt: pkt}
	cmdr.PushMsg(msg)
	pkt, _ = types.NewPacket(types.PktTUData, []byte("this is unsequenced packet 2"))
	msg = &types.Message{Msec: 5400, Pkt: pkt}
	cmdr.PushMsg(msg)
}

func main() {
	// parse args
	sAddrPtr := flag.String("server_addr", "127.0.0.1", "a string")
	sPortPtr := flag.String("server_port", "30010", "a string")
	cPortPtr := flag.String("client_port", "", "a string")
	flag.Parse()

	stopFlag := false
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt)
	go receive(sigC, &stopFlag)

	conn, err := net.NewKeepAliveTCPConn(*sAddrPtr, *sPortPtr, *cPortPtr)
	if err != nil {
		fmt.Println(err)
		return
	}
	cmdr := producer.NewCommander(conn.GetTube())
	produceCmds(cmdr)

	fmt.Println("Start sending messages")
	go cmdr.Order()
	go conn.Start()

	baseT := time.Now().UnixNano() / int64(time.Millisecond)
	curT := time.Now().UnixNano() / int64(time.Millisecond)
	for curT-baseT < 10000 && stopFlag == false {
		curT = time.Now().UnixNano() / int64(time.Millisecond)
	}
	err = conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Bye.")
}
