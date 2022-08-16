package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	//instructions
	io.WriteString(conn,
		"\r\nIN-MEMORY DATABASE\n"+
			"USE:\n"+
			"\tSET key value\n"+
			"\tGET key\n"+
			"\tDEL key \n"+
			"EXAMPLE:\r\n"+
			"\tSET fav chocolate\n"+
			"\tGET fav \n",
	)

	//read & write
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln) //split the string

		//logic
		if len(fs) < 1 {
			fmt.Fprintln(conn, "Please enter the command")
			continue
		}

		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\r\n", v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintf(conn, "EXPECTED VALUE\r\n")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND"+fs[1]+"\r\n")
			continue
		}
	}

}
