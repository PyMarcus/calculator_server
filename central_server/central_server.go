package central_server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/PyMarcus/calculator_server/interfaces"
	"github.com/PyMarcus/calculator_server/settings"
	"github.com/PyMarcus/calculator_server/sub_server"
	"github.com/PyMarcus/calculator_server/sum_server"
)

type CentralServer struct {
	Ip   string
	Port string
}

// New instance of SubServer class will be created
func New() *CentralServer {
	s := CentralServer{Ip: settings.GetCentralServerInfo("IP"), Port: settings.GetCentralServerInfo("PORT")}
	return &s
}

func manager(operatorInterface interfaces.IOperator){
	operatorInterface.Start()
}

// Start the server, by default on 127.0.0.1:9992
func (ss CentralServer) Start(){

	
	sum := sum_server.New()
	sub := sub_server.New()

	go manager(sum)
	go manager(sub)
	log.Printf("Running central server on %s:%s\n", ss.Ip, ss.Port)

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

func (ss CentralServer) handleConnection(conn net.Conn){
	var response string 

	for{

		buffer, _ := bufio.NewReader(conn).ReadString('\n')
	
		log.Println("[+] Received!")

		to := strings.Split(buffer, ",")[0]


		if to == "SUBTRACAO"{
			log.Println("[+]Client says ", buffer)
			response = ss.post(buffer, "sub")
		}else{
			log.Println("[+]Client says ", buffer)
			response = ss.post(buffer, "sum")
		}
		conn.Write([]byte(response + "\n"))
		log.Println("[+]Central Server: Sended response to client!")
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


	log.Printf("[+]Sending data %s to server %s\n", strings.TrimSpace(strings.Split(data, "\n")[0]), server)

	conn.Write([]byte(data))

	response, _ = bufio.NewReader(conn).ReadString('\n')
	return response
}
