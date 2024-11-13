package online

import (
	"encoding/json"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有源
		return true
	},
}

// 创建一个WsKlineEvent通道
var klineData chan *binance.WsKlineEvent

// 修改KlineBn函数，使其返回一个WsKlineEvent的通道
func KlineBn() {

	wsKlineHandler := func(event *binance.WsKlineEvent) {
		klineData <- event
		fmt.Println(<-klineData)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	// LTCBTC, 1m
	doneC, _, err := binance.WsKlineServe("BTCUSDT", "1m", wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(<-doneC)
}

// WebSocket处理函数
func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 添加CORS设置
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有来源
	// 升级HTTP连接到WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}
	defer conn.Close()

	// 向客户端发送欢迎消息
	welcomeMessage := []byte("Hello, this is a WebSocket server!")
	err = conn.WriteMessage(websocket.TextMessage, welcomeMessage)
	if err != nil {
		log.Println("Write:", err)
		return
	}
	// 增加一个循环来持续处理K线数据
	go func() {
		for {
			select {
			case event := <-klineData:
				// 将WsKlineEvent转换为JSON格式
				log.Println(event)
				eventJSON, err := json.Marshal(event)
				if err != nil {
					log.Println("Marshal event:", err)
				}
				err = conn.WriteMessage(websocket.TextMessage, eventJSON)
				if err != nil {
					log.Println("Write:", err)
					return
				}

			}
		}
	}()

	// 读取客户端发送的消息
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read:", err)
			break
		}
		log.Printf("Received message: %s\n", message)
		// 接收K线数据并推送给客户端

	}

}

func WsBn() {
	// 获取K线数据通道
	klineData = make(chan *binance.WsKlineEvent, 1000)
	go KlineBn()
	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(":8079", nil))
}
func DepthBn() {
	wsDepthHandler := func(event *binance.WsDepthEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, stopC, err := binance.WsDepthServe("LTCBTC", wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopC to exit
	go func() {
		time.Sleep(60 * time.Second)
		stopC <- struct{}{}
	}()
	// remove this if you do not want to be blocked here
	fmt.Println(<-doneC)
	//<-doneC
}
