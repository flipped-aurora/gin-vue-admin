package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
	"time"
)

func main() {
	// 连接到NATS服务器
	url := "nats://81.69.19.231:4222"
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	c := 10
	wg1 := &sync.WaitGroup{}
	wg2 := &sync.WaitGroup{}
	wg1.Add(c)
	wg2.Add(c)
	now := time.Now()
	// 创建一个订阅者
	sub, err := nc.Subscribe("greet.*", func(msg *nats.Msg) {
		name := msg.Subject[len("greet."):]
		msg.Respond([]byte("hello, " + name + string(msg.Data)))
		//fmt.Println(len(rCh))
		//fmt.Println(string(msg.Data))
		defer wg2.Done()
		//<-rCh
	})
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	for i := 0; i < c; i++ {
		// 发送请求并等待响应
		i := i
		go func() {
			defer wg1.Done()
			reply, err := nc.Request("greet.joe", []byte(fmt.Sprintf("%v", i)), time.Second*10)
			if err != nil {
				fmt.Println(err, i)
			}
			fmt.Println(string(reply.Data))
			//reply.Ack()
		}()
		//fmt.Printf("Received response: %s\n", string(reply.Data))
	}

	wg1.Wait()
	wg2.Wait()
	fmt.Println(time.Since(now))

}
