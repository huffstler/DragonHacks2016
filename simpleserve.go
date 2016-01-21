package main

import (
	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumbleutil"
	"fmt"
	"os"
)

func main() {
	// setting configuration
	defaultchannel := "DragonHacks2016"
	
	config := gumble.NewConfig()
	config.Username = "dhackstest"
	config.Address = "rlyshw.com:64738"
	config.TLSConfig.InsecureSkipVerify = true

	keepAlive := make(chan bool)
	// generatiing a client instance
	client := gumble.NewClient(config)

	client.Attach(gumbleutil.AutoBitrate)
	// attaching client 
	client.Attach(gumbleutil.Listener{
		TextMessage: func(e *gumble.TextMessageEvent) {
			fmt.Printf("Message received: %s\n", e.Message)
			e.Message = "joining"
			client.Send(e)
		},
	})

	// on-connect, perform the following
	client.Attach(gumbleutil.Listener{
		Connect: func(e *gumble.ConnectEvent) {
			fmt.Printf("Connected!!: %s\n", e.MaximumBitrate)
			fmt.Printf("User Channel: %s\n", client.Self.Channel.Name)
			for i, _ := range client.Channels {
				if(client.Channels[i].Name == defaultchannel) {
					client.Self.Move(client.Channels[i])
					break;
				}
			}
			
			// outlining for audio operations
		},
	})


	client.Attach(gumbleutil.Listener{
		Disconnect: func(e *gumble.DisconnectEvent) {
			keepAlive <- true
		},
	})

	
	// connecting the client
	if err := client.Connect(); err != nil {
		panic(err)
		os.Exit(1)
	}
	// concurrency???
	<-keepAlive
}
