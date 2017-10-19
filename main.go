package main

import (
	"fmt"
	"net"
	"os"
	"errors"
	"strings"
)


func main() {
	l, err := net.Listen("tcp", ":8080")
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

	request := string(buff[:length])
//	func SplitN(s, sep string, n int) []string
	requestParts := strings.SplitN(request, "\r\n\r\n", 2)

	if len(requestParts) != 2 {
		return errors.New("400")
	}
	headerLines := strings.Split(requestParts[0], "\r\n")
	method := strings.Split(headerLines[0], " ")

	headers := make(map[string]string)
	for _, line := range headerLines[1:] {
		lineParts := strings.SplitN(line, ":", 2)
		if len(lineParts) != 2 {
			return errors.New("400")
		}
		headers[strings.TrimSpace(lineParts[0])] = lineParts[1]
	}

	response := "HTTP/1.0 200 ok\r\n"
	response += "Content-type: text/html\r\n\r\n"
	response += "<pre>"
	response += "\n\nPART 1:\n\n" + requestParts[0]
	response += "\n\nPART 2:\n\n" + requestParts[1]
	response += "\n\n</pre>\n"
//	response += "hello\n" + headers[1]
//	response += "\n\n mello" + met[0]
	if method[0] == "GET" {
		response += "This request method is: " + method[0] + ", it has no body\n"
	}
	if method[0] == "POST" {
		response += "This request method is: " + method[0] + ", it has body\n"
	}
	response += "\n\nPARSED HEADERS:\n\n" + fmt.Sprintf("%+v", headers)

	conn.Write([]byte(response))

	conn.Close()

	return nil
}

