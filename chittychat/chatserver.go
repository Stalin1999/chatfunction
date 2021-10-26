package chittychat

import (
	"log"
	"sync"
	"time"
	"math/rand"

	//chittychat "github.com/Stalin1999/chatfunction/chittychat"
)

type messageUnit struct{
	ClientName string
	MessageMessage string
	MessageUniqueCode int
	ClientUniqueCode int
}

type messageHandle struct {
	MessageQueue []messageUnit
	mu sync.Mutex
}

var messageHandleObject = messageHandle{}

type ChatServer struct {
	//chittychat.UnimplementedServiceServer;
}



func (is *ChatServer) ChatService(csi serviceChatServiceServer) error {


	ClientUniqueCode := rand.Intn(1e6)
	errch := make(chan error)

	//Recive messages 
	go receiveFromStream(csi, ClientUniqueCode, errch)

	//Send messages
	go sendToStream(csi, ClientUniqueCode, errch)

	return <- errch

}

func receiveFromStream(csi_ serviceChatServiceServer, clientUniqueCode_ int, errch_ chan error){
	for {
		mssg, err := csi_.Recv()
		if err!=nil{
			log.Printf("Error in receiving from the client :: %v", err)
			errch_ <- err
		} else {
		messageHandleObject.mu.Lock()
		messageHandleObject.MessageQueue = append(messageHandleObject.MessageQueue, messageUnit{
			ClientName : mssg.User,
			MessageMessage : mssg.Message,
			MessageUniqueCode : rand.Intn(1e8),
			ClientUniqueCode: clientUniqueCode_,
		})
		
		messageHandleObject.mu.Unlock()
		
		log.Printf("&v", messageHandleObject.MessageQueue[len(messageHandleObject.MessageQueue)-1])
		}
	}
}

func sendToStream(csi_ serviceChatServiceServer, clientUniqueCode_ int, errch_ chan error){
	for {
		for {
			time.Sleep(500 * time.Millisecond)
			
			messageHandleObject.mu.Lock()

			if len(messageHandleObject.MessageQueue) == 0 {
				messageHandleObject.mu.Unlock()
				break
			}
			
			senderUniqueCode := messageHandleObject.MessageQueue[0].ClientUniqueCode
			senderName4Client := messageHandleObject.MessageQueue[0].ClientName
			message4Client := messageHandleObject.MessageQueue[0].MessageMessage

			if senderUniqueCode != clientUniqueCode_ {
				err := csi_.Send(&Broadcast{User:senderName4Client,Message:message4Client})

				if err != nil{
					errch_ <- err
				}

				messageHandleObject.mu.Lock()
				
				if len(messageHandleObject.MessageQueue) > 1 {
					messageHandleObject.MessageQueue = messageHandleObject.MessageQueue[1:]
				} else {
					messageHandleObject.MessageQueue = []messageUnit{}
				}

				messageHandleObject.mu.Unlock()
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}