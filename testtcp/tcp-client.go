package main

import (
	//"bufio"
	"flag"
	"fmt"
	"net"
	"time"
)

var (
	serverAddr string
	readData   bool
	writeData  bool
	interval   int
)

func init() {
	flag.StringVar(&serverAddr, "s", "127.0.0.1:4444", "send data to server")
	flag.BoolVar(&readData, "r", false, "open read data")
	flag.BoolVar(&writeData, "w", false, "open write data")
	flag.IntVar(&interval, "i", 1000*1000*1000, "1(second) = 1000 * 1000 * 1000(nanosecond)")
	flag.Parse()
}

func main() {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		if writeData || readData {
			if writeData {
				fmt.Println("start write")
				n, err := conn.Write([]byte("0123456789"))
				if err != nil {
					fmt.Println("write err", err)
				}
				fmt.Println("send", n, "byte to", conn.RemoteAddr())
			}
			if readData {
				fmt.Println("start read")
				buf := make([]byte, 10)
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println("write err", err)
				}
				fmt.Println("read", n, "byte:", string(buf), "from", conn.RemoteAddr())
			}
		} else {
			fmt.Println("waiting", interval, "nanosecond")
		}

		time.Sleep(time.Duration(interval) * time.Nanosecond)
	}
}
