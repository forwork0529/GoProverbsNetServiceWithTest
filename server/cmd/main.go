package main

import (
	"log"
	"math/rand"
	"net"
	"os"
	"prvbNetServ/packages/conHandlers"
	"prvbNetServ/packages/getFrom"
	"time"
)

var (
	maxUsers int = 3
	network = "tcp4"
	addr = "127.0.0.1:12345"

	ch = make(chan struct{}, maxUsers)

)

func main(){
	rand.Seed(time.Now().UnixNano())
	data := getFrom.Files()
	logs := newLogger()
	l, err :=  net.Listen(network, addr)
	if err != nil{
		log.Fatal(err)
	}
	defer l.Close()
	logs.Println("server started")
	for{
		ch <- struct{}{}
		con, err := l.Accept()
		if err != nil{
			log.Fatal(err)
		}
		logs.Println("клиент установил соединение")
		go func (){
			defer func(){
				//Завершение соединения сервером в коде отсутсвует по заданию !!!
				/*
				_ = con.Close()
				logs.Println("клиент завершил соединение")
				<- ch // по заданию я не зыкрываю соединение , значит нельзя запустить больше 3 соединений
				*/
			}()
			err = conHandlers.ProverbsHandler(con, conHandlers.Params( data, logs, 10))
			if err != nil{
				logs.Println(err)
			}

		}()

	}
}

func newLogger()*log.Logger{
	output, err := os.Create("files/logs.txt")
	if err != nil{
		log.Fatal(err)
	}
	return log.New(output,"INFO\t",log.Ldate|log.Ltime)
}
