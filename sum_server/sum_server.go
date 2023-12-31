package sum_server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"strconv"
	"github.com/PyMarcus/calculator_server/settings"
)

type SumServer struct {
	Ip   string
	Port string
}

// New instance of SubServer class will be created
func New() *SumServer {
	s := SumServer{Ip: settings.GetSumServerInfo("IP"), Port: settings.GetSumServerInfo("PORT")}
	return &s
}

// Solve operations like x-y, with N parameters
func (ss SumServer)solve(operands []string) int{
	sum := 0

	log.Println("[+]Operands to solving: ", operands)

	for _, num := range operands{
		n, _ := strconv.Atoi(strings.TrimSpace(num))
		sum += n
	}
	log.Println("[+]Result of sum: ", sum)
	return sum
}

// Start the server, by default on 127.0.0.1:9991
func (ss SumServer) Start(){

	log.Printf("Running sum server on %s:%s\n", ss.Ip, ss.Port)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", ss.Ip, ss.Port))

	if err != nil{
		log.Fatalln("Error to connect ", err)
	}

	for{
		conn, _ := listener.Accept()
		defer conn.Close()

		log.Println("[+]Connection received from: ", conn.RemoteAddr())

		go ss.handleConnection(conn)
	}
}

func (ss SumServer) handleConnection(conn net.Conn){
	operands := []string{} 
	var response int 
	buffer, _ := bufio.NewReader(conn).ReadString('\n')

	parserBuffer := strings.Split(strings.TrimSpace(buffer), ",")[1:]
	operands = append(operands, parserBuffer...)

	response = ss.solve(operands)

	conn.Write([]byte(strconv.Itoa(response) + "\n"))
	log.Println("[+]Sum Server says: 'Sended response to central server!'")	
}
