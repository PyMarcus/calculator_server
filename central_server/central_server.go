package central_server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"github.com/PyMarcus/calculator_server/settings"
	"github.com/PyMarcus/calculator_server/sub_server"
	"github.com/PyMarcus/calculator_server/sum_server"
	"github.com/PyMarcus/calculator_server/interfaces"
)

type CentralServer struct {
	Ip   string
	Port string
}

// New instance of SubServer class will be created
func (ss CentralServer) New() *CentralServer {
	s := CentralServer{Ip: settings.GetCentralServerInfo("IP"), Port: settings.GetCentralServerInfo("PORT")}
	return &s
}

func manager(operatorInterface interfaces.IOperator){
	operatorInterface.Start()
}

// Start the server, by default on 127.0.0.1:9992
func (ss CentralServer) Start(){
	server := ss.New()

	log.Printf("Running central server on %s:%s\n", server.Ip, server.Port)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.Ip, server.Port))

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

func (ss CentralServer) handleConnection(conn net.Conn){
	var response string 

	sum := sum_server.New()
	sub := sub_server.New()

	go manager(sum)
	go manager(sub)

	for{
		buffer, _ := bufio.NewReader(conn).ReadString('\n')

		to := strings.Split(buffer, ",")[0]


		if to == "SUBTRACAO"{
			response = ss.post(buffer, "sub")
		}else{
			response = ss.post(buffer, "sum")
		}
		conn.Write([]byte(response))
		log.Println("Sended!")
	}
}

// post act like a client to send data to sum or sub servers
func (ss CentralServer) post(data string, server string) (response string){
	var conn net.Conn 
	var err error 

	if server == "sub"{
		conn, err = net.Dial("tcp", fmt.Sprintf("%s:%s", settings.GetSubServerInfo("IP"), settings.GetSubServerInfo("PORT")))
	}else{
		conn, err = net.Dial("tcp", fmt.Sprintf("%s:%s", settings.GetSumServerInfo("IP"), settings.GetSumServerInfo("PORT")))
	}
	defer conn.Close()

	if err != nil{
		log.Fatalln(err)
	}

	log.Printf("Sending data %s to server %s\n", data, server)

	conn.Write([]byte(data))

	response, _ = bufio.NewReader(conn).ReadString('\n')
	return response
}
