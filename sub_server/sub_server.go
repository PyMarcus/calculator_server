package sub_server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"strconv"
	"github.com/PyMarcus/calculator_server/settings"
)

type SubServer struct {
	Ip   string
	Port string
}

// New instance of SubServer class will be created
func New() *SubServer {
	s := SubServer{Ip: settings.GetSubServerInfo("IP"), Port: settings.GetSubServerInfo("PORT")}
	return &s
}

// Solve operations like x-y, with N parameters
func (ss SubServer)solve(operands []string) int{
	sub := 0
	intOperands := make([]int, len(operands))

	log.Println("Operands to solving: ", operands)
	for index, num := range operands{
		n, _ := strconv.Atoi(strings.TrimSpace(num))
		if index == 0{
			intOperands[index] = n
		}else{
			intOperands[index] = -n
		}
	}
	
	for _, num := range intOperands{
		sub += num
	}
	log.Println("Result: ", sub)
	return sub
}

// Start the server, by default on 127.0.0.1:9990
func (ss SubServer) Start(){

	log.Printf("Running sub server on %s:%s\n", ss.Ip, ss.Port)

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

func (ss SubServer) handleConnection(conn net.Conn){
	operands := []string{} 
	var response int 
	buffer, _ := bufio.NewReader(conn).ReadString('\n')

	parserBuffer := strings.Split(strings.TrimSpace(buffer), ",")[1:]
	operands = append(operands, parserBuffer...)

	response = ss.solve(operands)

	conn.Write([]byte(strconv.Itoa(response) + "\n"))
	log.Println("[+]Sub Server says: 'Sended response to central server!'")
	
}
