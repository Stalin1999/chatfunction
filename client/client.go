package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
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

	for {
		reader := bufio.NewReader(os.Stdin)
		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read message", err)
		}
		SendRequest(c, clientMessage)
	}
}

type clienthandle struct {
    stream chittychat.Service_ChatServiceClient
    clientName string
}

func SendRequest(c chittychat.ServiceClient, message2 string) {
	// Between the curly brackets are nothing, because the .proto file expects no input.
	message := chittychat.Publish{}

	response, err := c.ChatService(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling: %s", err)
	}

	fmt.Printf("Response from the Server: %s \n", response)
}

func (ch *clienthandle) recieveMessage(){
    for {
        mssg, err := ch.stream.Recv()
        if err != nil {
            log.Printf("Error recieving message from server", err)
        }

        fmt.Println(mssg.User, mssg.Message)
    }
}
