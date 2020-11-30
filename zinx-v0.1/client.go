package main

import (
	"fmt"
	"io"
	"net"

	"github.com/lumigogogo/zinx/znet"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("conn error!")
		return
	}

	// readData := make([]byte, 512)
	// conn.Write([]byte("hello"))

	go func() {
		for {
			// _, err = conn.Read(readData)
			// if err != nil {
			// 	return
			// }

			//封装一个msg1包
			msg1 := znet.NewMessage(1, 5, []byte{'h', 'e', 'l', 'l', 'o'})
			// msg1 := znet.NewMessage(5, 1, []byte{'h', 'e', 'l', 'l', 'o'})

			sendData1, err := znet.Pack(msg1)
			if err != nil {
				fmt.Println("client pack msg1 err:", err)
				return
			}

			msg2 := znet.NewMessage(2, 7, []byte{'w', 'o', 'r', 'l', 'd', '!', '!'})
			// msg2 := znet.NewMessage(7, 2, []byte{'w', 'o', 'r', 'l', 'd', '!', '!'})

			sendData2, err := znet.Pack(msg2)
			if err != nil {
				fmt.Println("client temp msg2 err:", err)
				return
			}

			//将sendData1，和 sendData2 拼接一起，组成粘包
			sendData1 = append(sendData1, sendData2...)

			//向服务器端写数据
			conn.Write(sendData1)
		}
	}()

	go func() {
		for {
			//1 先读出流中的head部分
			headData := make([]byte, znet.DataHeadLen)
			_, err := io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
			if err != nil {
				fmt.Println("read head error")
			}
			// fmt.Println("headData: ", headData)
			//将headData字节流 拆包到msg中
			msgHead, err := znet.Unpack(headData)
			if err != nil {
				fmt.Println("server unpack err:", err)
				return
			}

			if msgHead.GetDataLen() > 0 {
				//msg 是有data数据的，需要再次读取data数据
				msg := msgHead.(*znet.Message)
				msg.SetData(make([]byte, msg.GetDataLen()))

				//根据dataLen从io中读取字节流
				_, err := io.ReadFull(conn, msg.GetData())
				if err != nil {
					fmt.Println("server unpack data err:", err)
					return
				}

				fmt.Println("==> Recv Msg: ID=", msg.GetMsgID(), ", len=", msg.GetDataLen(), ", data=", string(msg.GetData()))
			}
		}
	}()

	select {
	// case <-time.After(10 * time.Second):
	// 	return
	}
}
