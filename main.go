package main

import (
	"fmt"
	"net"
	"os"
	"errors"
)


func main() {
	l, err := net.Listen("tcp", ":65535")
	if err != nil {
		fmt.Println(err)
		os.Exit(111)
	}

	for true {
		err = huificate(l)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	fmt.Println("end")
}

func huificate(listener net.Listener) error {
	conn, err := listener.Accept()
	if err != nil {
		return err
	}

	buff := make([]byte, 1024)
	length, err := conn.Read(buff)
	if err != nil {
		return err
	}

	word := string(buff[:length])
	if word == "stop\n" {
		return errors.New("stop word received")
	}

	conn.Write([]byte("hu" + word))

	conn.Close()

	return nil
}
