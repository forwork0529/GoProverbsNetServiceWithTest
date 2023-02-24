package conHandlers

import (
	"bufio"
	"log"
	"net"
	"strings"
	"sync"
	"testing"
)

func TestProverbsHandler(t *testing.T) {
	logs := log.Default()
	data := [][]byte{[]byte("Cgo is not Go.\n")}
	srv ,cl := net.Pipe()
	var errSrv error

	wg := sync.WaitGroup{}
	wg.Add(1)


	go func(srv net.Conn, data [][]byte){
		errSrv = ProverbsHandler(srv, &params{
			data,logs, randIntTest, waitTest, 1} )
		_ = srv.Close()
		if errSrv != nil{
			t.Error(errSrv)
		}
		wg.Done()
	}(srv, data)


	_, _ = cl.Write([]byte("proverbs\n"))
	r := bufio.NewReader(cl)
	msg, _ := r.ReadString('\n')
	msg = strings.Trim(msg,"\n")
	msg = strings.Trim(msg,"\r")

	if msg != "Cgo is not Go."{
		t.Errorf("answer: %v not expected", msg)
	}
	logs.Println("got = want")
	wg.Wait()
	_ = cl.Close()
}