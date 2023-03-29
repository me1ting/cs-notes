package main

import (
	"fmt"
	"net"
	"strconv"
)

func clientRun(c *CountdownLatch) {
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:1099")
	conn, _ := net.DialUDP("udp", nil, addr)
	defer conn.Close()

	for i := 0; i < 10; i++ {
		n, oobn, _ := conn.WriteMsgUDP([]byte(""+strconv.Itoa(i)), nil, nil)
		fmt.Printf("send: %v, return: %v, %v\n", i, n, oobn)
	}
	c.CountDown()
}

func serverRun(c *CountdownLatch) {
	conn, _ := net.ListenPacket("udp", ":1099")
	defer conn.Close()

	buf := make([]byte, 10)
	for i := 0; i < 10; i++ {
		n, addr, _ := conn.ReadFrom(buf)
		fmt.Printf("recive: %v, return: %v, %v\n", string(buf[0:n]), n, addr)
	}

	c.CountDown()
}

func main() {
	var c = NewCountDownLatch(2)

	go clientRun(&c)
	go serverRun(&c)

	c.Await()
}

type CountdownLatch struct {
	countDown chan int
	await     chan bool
	count     int
}

func NewCountDownLatch(count int) CountdownLatch {
	var c = CountdownLatch{
		make(chan int, count),
		make(chan bool),
		count,
	}

	go func() {
		count := 0
		for i := range c.countDown {
			count += i
			if count == c.count {
				break
			}
		}
		close(c.await)
	}()

	return c
}

func (c *CountdownLatch) CountDown() {
	c.countDown <- 1
}

func (c *CountdownLatch) Await() {
	<-c.await
}
