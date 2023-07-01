package main

import (
	"flag"
	"log"
	"github.com/PyMarcus/calculator_server/central_server"
	"github.com/PyMarcus/calculator_server/client"
)

func run(){
	option := flag.String("start", "server", "Choice a option to running. Example: cmd [-server/-client]")
	flag.Parse()
	switch(*option){
	case "server":
		log.Println("[V]Server")
		cs := central_server.New() 
		cs.Start()
		break
	case "client":
		log.Println("[V]Client")
		client.Request()
		break
	default:
		log.Println("Invalid option. Choice a option to running. Example: cmd [-server/-client]")
	}
}

func main() {
	run()
}
