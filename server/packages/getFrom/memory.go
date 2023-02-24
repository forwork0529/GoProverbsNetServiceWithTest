package getFrom

import (
	"bytes"
	"log"
	"os"
)

func Files()[][]byte{
	b, err := os.ReadFile("files/input.txt")
	if err != nil{
		log.Fatal(err)
	}
	return  bytes.Split(b, []byte("\n"))
}

