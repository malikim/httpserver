package main

import (
	"fmt"
	"net"
	"os"
	"errors"
	"strings"
	"net/url"
	"encoding/json"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(111)
	}

	for true {
		err = handler(l)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	fmt.Println("end")
}

func handler(listener net.Listener) error {
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

	requestParts := strings.SplitN(request, "\r\n\r\n", 2)

	if len(requestParts) != 2 {
		return errors.New("400")
	}
	headerLines := strings.Split(requestParts[0], "\r\n")
	method := strings.Split(headerLines[0], " ")
	headers, err := parseHeaders(headerLines[1:])
	if err != nil {
		return err
	}


	urlParse(method[1])

	response := "HTTP/1.0 200 ok\r\n"
	response += "Content-type: text/html\r\n\r\n"
	response += "<pre>"
	response += "\n\nPART 1:\n\n" + requestParts[0]
	response += "\n\nPART 2:\n\n" + requestParts[1]
	response += "</pre>"


	var lapk map[string]string
	if method[0] == "GET" {
		response += "This request method is: " + method[0] + ", it has no body\r\n\r\n"
	} else if method[0] == "POST" {
		response += "This request method is: " + method[0] + ", it has body\n"
		contentType := headers["content-type"]
		if strings.Contains(contentType, "application/x-www-form-urlencoded") {
			response += "This request's body is urlencoded\n"
			lapk, err = formUrlParse(requestParts[1])
			response += fmt.Sprintf("Your urlencoded kirillic name is: %s %s \n", lapk["Ivan"], lapk["Ivanov"])
		} else if strings.Contains(contentType, "application/json") {
			response += "This request's body type is json\n"
			lapk, err = jsonParse(requestParts[1])
			response += fmt.Sprintf("Your name is: %s %s \n", lapk["firstName"], lapk["lastName"])
			firstName := lapk["firstName"]
				if strings.Contains(firstName, "firstName") {
					response += "Your first name is: " + lapk["firstName"]
				}
			if err != nil {
				return err
			}
			response += "Parsed json body:\n\n" + fmt.Sprintf("%+v", lapk)


		} else if strings.Contains(contentType, "multipart/form-data") {
			response += "This request's body type is multipart\n"
			lapk, err = multipartParse(requestParts[1])
			response += fmt.Sprintf("Your multipart name is: %s %s \n", lapk["Ivan"], lapk["Ivanov"])
		}
	}


	if method[0] == "GET" {
		response += "\r\nPARSED HEADERS:\n\n" + fmt.Sprintf("%+v", headers)
	} else if method[0] == "POST" {
		response += "\r\nPARSED HEADERS:\n\n" + fmt.Sprintf("%+v", headers)
	}

	conn.Write([]byte(response))

	conn.Close()

	return nil
}

func parseHeaders(lines []string) (map[string]string, error) {
	headers := make(map[string]string)
	for _, line := range lines {
		lineParts := strings.SplitN(line, ":", 2)
		if len(lineParts) != 2 {
			return nil, errors.New("400")
		}
		headerName := strings.ToLower(strings.TrimSpace(lineParts[0]))
		headers[headerName] = strings.TrimSpace(lineParts[1])
	}
	return headers, nil


}

func formUrlParse(urlStr []string) (map[string]string, error) {
	data := strings.Split(urlStr[0], "&")
	dataResult := make(map[string]string)
	for _, str := range data {
		strParts := strings.Split(str, "=")

		if len(strParts) != 2 {
			return nil, errors.New("400")
		}
		dataResult[strParts[0]] = strParts[1]
	}

	return dataResult, nil
}
func urlParse(urlStr string) (*url.URL, error) {
	// вкурить и сделать без функции
	u, err := url.Parse(urlStr)
	if err == nil {
		fmt.Println(u.Path)
		fmt.Println(u.RawPath)
		fmt.Println(u.String())
	}
	return u, err


}

//func Marshal(v interface{}) ([]byte, error)
//requestParts[1]
func jsonParse(jsonStr string) (map[string]string, error) {
	//	var dat map[string]interface{}
	var data map[string]string
	err := json.Unmarshal([]byte(jsonStr), &data)
	return data, err
}

func multipartParse(mulripartStr string) (map[string]string, error) {
	var data map[string]string
	err :=
	return data, err
}
//lapk
/*
func jsonResponse (value map[string]string) ([]byte, error) {
	if strings.Contains(value, "firstName") {
		response +=
	}
	return a, b
}




if strings.Contains(contentType, "application/json") {
			response += "This request's body type is json\n"
			lapk, err = jsonParse(requestParts[1])
			if err != nil {
				return err
for _, line := range lines {
		lineParts := strings.SplitN(line, ":", 2)
		if len(lineParts) != 2 {
			return nil, errors.New("400")
		}
		headerName := strings.ToLower(strings.TrimSpace(lineParts[0]))
		headers[headerName] = strings.TrimSpace(lineParts[1])
	}
 */
