package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
    "strings"

	//"os/user"

	//"google.golang.org/genproto/googleapis/streetview/publish/v1"
	"google.golang.org/grpc"

	//chatserver "github.com/Stalin1999/chatfunction/chittychat"
	chittychat "github.com/Stalin1999/chatfunction/chittychat"
)

func main() {
	// Create a virtual RPC Client Connection on port 9080 WithInsecure (because of http)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	// Defer means: When this function returns, call this method (meaning, one main is done, close connection)
	defer conn.Close()

	// Create new Client from generated gRPC code from proto
	c := chittychat.NewServiceClient(conn)

	stream, err := c.ChatService(context.Background())
	if err != nil {
		log.Fatalf("Failed to call ChatService :: %v", err)
	}

    ch := clienthandle{stream: stream}
    ch.clientConfig()
    go ch.sendMessage()
    go ch.recieveMessage()

    bl := make(chan bool)
    <-bl
}

type clienthandle struct {
    stream chittychat.Service_ChatServiceClient
    clientName string
}

func(ch *clienthandle) clientConfig(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Your UserName: ")
	Name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf(" Failed to read from console: &v", err)
	}
	ch.clientName = strings.Trim(Name,"\r\n")
}

func (ch *clienthandle) sendMessage() {
    for {
        
        reader := bufio.NewReader(os.Stdin)
        clientMessage, err := reader.ReadString('\n')
        if err != nil {
            log.Fatalf("Failed to read message", err)
        }
        clientMessage = strings.Trim(clientMessage, "\r\n")

        clientMessageBox := &chittychat.Publish{
            User: ch.clientName,
            Message: clientMessage,
			Time: 0,
        }

        err = ch.stream.Send(clientMessageBox)
        
        if err != nil {
            log.Printf("Error sending the message to the server: %v", err)
        }
    }
}

func (ch *clienthandle) recieveMessage(){
    for {
        //mssg, err := ch.stream.Recv()
        //if err != nil {
        //    log.Printf("Error recieving message from server %s", err)
        //}

        //fmt.Printf("&s : &s : &s \n", mssg.User, mssg.Message, mssg.Time)

        fmt.Printf("Test")
    }
}
