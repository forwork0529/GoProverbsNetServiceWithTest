package conHandlers

import (
	"bufio"
	"errors"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

var (
	r, r2 *bufio.Reader
	cantRecognize error = errors.New("cant recognize command in the message")
	noEntryData = errors.New("no entry data")

	// mu sync.Mutex  // Seed, unlike the Rand.Seed method, is safe for concurrent use.
)


type randFunc func(int)int

type waitFunc func()

type params struct{
	data [][]byte
	logs *log.Logger
	rF randFunc
	wF waitFunc
	repeats int
}

func Params(data [][]byte, logs *log.Logger, repeats int)*params{
	return &params{data : data, logs : logs, rF : randInt, wF : wait, repeats : repeats}
}

func ProverbsHandler(conn net.Conn, par *params )error{

	if len(par.data) < 1{
		par.logs.Print(noEntryData.Error())
		return noEntryData
	}

	r = bufio.NewReader(conn)
	b, err := r.ReadBytes('\n')
	if err != nil{
		return err
	}

	msg := strings.TrimSuffix(string(b),"\n")
	msg = strings.TrimSuffix(msg, "\r")
	par.logs.Printf("get message from client: %v\n", msg)

	if msg != "proverbs"{

		_, err = conn.Write([]byte(cantRecognize.Error()))
		return err
	}

	for i := 0 ; i < par.repeats ; i ++{
		msg := par.data[par.rF(len(par.data))]
		msg = append(msg, []byte("\n")...)
		_, err = conn.Write(msg)
		if err != nil{
			return err
		}
		par.wF()
		par.logs.Printf( "Sent to client: "+string(msg))
	}
	return nil
}



func randInt(lenAr int)int{
	// mu.Lock()     // Seed, unlike the Rand.Seed method, is safe for concurrent use.
	res := rand.Intn(lenAr)
	// mu.Unlock()   // Seed, unlike the Rand.Seed method, is safe for concurrent use.
	return res
}

func randIntTest(lenAr int)int{
	return 0
}


func wait(){
	time.Sleep(time.Second * 3)
}

func waitTest(){
}



