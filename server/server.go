package server

import (
	"log"
	"fmt"
	"bufio"
	"os"
	"net"
)


const port = "8080"

//RunHost takes an ip as parameter and listen for tcp connections
func RunHost(ip string){
	ipAndPort := ip + ":" + port

	l, err := net.Listen("tcp", ipAndPort)

	if err != nil{
		log.Fatal("Error listen port: ", err)
	}

	fmt.Println("Listening on", ipAndPort)
	conn, err := l.Accept()

	if err != nil{
		log.Fatal("Error accepting connection: ", err)
	}

	fmt.Println("New connection accepted")

	for{
		hostHandler(conn)
	}
}

//hostHandler is a helper function that sends and receives messages
func hostHandler(conn net.Conn){
	reader := bufio.NewReader(conn)
	msg , err := reader.ReadString('\n') //return key

	if err != nil{
		log.Fatal("Error reading string: ", err)
	}

	fmt.Println("Message received: ", msg)

	fmt.Print("Send message: ")
	replyReader := bufio.NewReader(os.Stdin)
	replyMsg, err := replyReader.ReadString('\n')

	if err != nil{
		log.Fatal("Error replying messsage: ", err)
	}

	fmt.Fprint(conn, replyMsg)
}

//guestHandler is a helper function that sends and receives messages
func guestHandler(conn net.Conn){
	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)
	msg, err := reader.ReadString('\n')

	if err != nil{
		log.Fatal("Error reading string", err)
	}

	fmt.Fprint(conn, msg)

	replyReader := bufio.NewReader(conn)
	replyMsg, err := replyReader.ReadString('\n')

	if err != nil{
		log.Fatal("Error reading string: ", err)
	}

	fmt.Println("Message received: ", replyMsg)
}

//RunGuest takes an ip as parameter and connects to the host.
func RunGuest(ip string){
	ipAndPort := ip + ":" + port
	conn, err := net.Dial("tcp", ipAndPort)

	if err != nil{
		log.Fatal("Error trying to connect to the host: ", err)
	}

	for{
		guestHandler(conn)
	}

}
