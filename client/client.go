package client

import (
	"os"
	"bufio"
	"log"
	"net"
	"github.com/PyMarcus/calculator_server/settings"
	"fmt"
)

// Request the service on 127.0.0.1:9992
func Request(){

	centralServiceIp := settings.GetCentralServerInfo("IP")
	centralServicePort := settings.GetCentralServerInfo("PORT")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", centralServiceIp, centralServicePort))
	defer conn.Close()

	if err != nil{
		log.Fatalln(err)
	}

	log.Printf("[+]Connected into Central server on %s:%s\n", centralServiceIp, centralServicePort)

	for{
		data := input()
		if data == "exit"{
			conn.Close()
			break
		}

		log.Println("Sending message: ", data)
		conn.Write([]byte(data + "\n"))

		buffer, _ := bufio.NewReader(conn).ReadString('\n')

		log.Println("[+] Server response: ", buffer)
	}
}

func input() string{

	in := bufio.NewScanner(os.Stdin)

	in.Scan()

	scanner := in.Text()

	return scanner 
}
