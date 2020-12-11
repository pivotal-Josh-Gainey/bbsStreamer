package main

import (
	"code.cloudfoundry.org/bbs"
	"code.cloudfoundry.org/bbs/events"
	"code.cloudfoundry.org/lager"
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 4 {
		fmt.Println("Usage:", os.Args[0], "<ca-crt> <client-crt> <client-key>")
		return
	}
	ca := os.Args[1]
	crt := os.Args[2]
	key := os.Args[3]

	logger := lager.NewLogger("BBS-STREAMER")


//	client, err := bbs.NewClient("https://bbs.service.cf.internal:8889", "ca.crt", "crt.crt", "key.key", 0, 0)
	client, err := bbs.NewClient("https://bbs.service.cf.internal:8889", ca, crt, key, 0, 0)

	if err != nil {
		log.Printf("failed to subscribe to lrp events: " + err.Error())
	}

	eventSource, err := client.SubscribeToEvents(logger)
	if err != nil {
		log.Printf("failed to subscribe to lrp events: " + err.Error())
	}

	subscriptionChan := make(chan events.EventSource, 1)

	var subscription events.EventSource

	subscriptionChan <- eventSource

	for {
		select {
		case subscription = <-subscriptionChan:
			if subscription != nil {
				for{
					event, _ := subscription.Next()
					fmt.Println(event.String())
				}
			}
		}
	}
}