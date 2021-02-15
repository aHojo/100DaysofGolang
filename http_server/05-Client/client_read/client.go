package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// bs, err := ioutil.ReadAll(conn)
	// if err != nil {
	// 	log.Println(err)
	// }

	// fmt.Println(string(bs))

	connbuf := bufio.NewReader(conn)
	for {
		str, err := connbuf.ReadString('\n')
		if err != nil {
			break
		}

		if len(str) > 0 {
			fmt.Println(str)
		}
	}
}
