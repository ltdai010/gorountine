package main

import (
	"fmt"
)

var (
	quit = make(chan int)
)

func main() {
	listSubscriber := make(map[string]*Subscriber)
	listTopic := make(map[string]*Topic)
	listPublisher := make(map[string]*Publisher)
	//create topic
	listTopic["test"] = &Topic{name: "test"}
	top := listTopic["test"]
	top.init()
	//crate 2 subscriber
	listSubscriber["Dai"] = &Subscriber{name: "Dai"}
	listSubscriber["Diaz"] = &Subscriber{name: "Diaz"}
	sub1 := listSubscriber["Dai"]
	sub1.init()
	sub1.subscribe(top)
	sub2 := listSubscriber["Diaz"]
	sub2.init()
	sub2.subscribe(top)
	//create a publisher
	listPublisher["Dia"] = &Publisher{name: "Dia"}
	pub1 := listPublisher["Dia"]
	listPublisher["Dat"] = &Publisher{name: "Dat"}
	pub2 := listPublisher["Dat"]
	publishToATopic(pub1, top, listSubscriber, "pub 1 post")
	publishToATopic(pub2, top, listSubscriber, "pub 2 post")
}

func publishToATopic(pub *Publisher, top *Topic, listSubscriber map[string]*Subscriber, content string) {
	//publish a content
	go pub.publish(top, content)
	//print notice of this subscriber
	go printNotice(listSubscriber, top.subscriber)
	//send to all subscriber
	for {
		if top.notifyAll(listSubscriber) {
			return
		}
	}
}



func printNotice(subscriberList map[string]*Subscriber, subscriberName []string)  {
	for _, i := range subscriberName {
		fmt.Println(<-subscriberList[i].receiver + " to " + i)
	}
	quit <- 0
}





