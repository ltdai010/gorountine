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
	//create 2 topic
	listTopic["test"] = &Topic{name: "test"}
	top1 := listTopic["test"]
	top1.init()
	listTopic["test1"] = &Topic{name: "test1"}
	top2 := listTopic["test1"]
	top2.init()
	//crate 3 subscriber
	listSubscriber["Dai"] = &Subscriber{name: "Dai"}
	listSubscriber["Diaz"] = &Subscriber{name: "Diaz"}
	listSubscriber["Hung"] = &Subscriber{name: "Hung"}
	sub1 := listSubscriber["Dai"]
	sub1.init()
	sub1.subscribe(top1)
	sub2 := listSubscriber["Diaz"]
	sub2.init()
	sub2.subscribe(top1)
	sub2.subscribe(top2)
	sub3 := listSubscriber["Hung"]
	sub3.init()
	sub3.subscribe(top2)
	//create 2 publisher
	listPublisher["Dia"] = &Publisher{name: "Dia"}
	pub1 := listPublisher["Dia"]
	listPublisher["Dat"] = &Publisher{name: "Dat"}
	pub2 := listPublisher["Dat"]
	publishToATopic(pub1, top1, listSubscriber, "pub 1 post")
	publishToATopic(pub2, top2, listSubscriber, "pub 2 post")
}

//this method is used for a publisher to post to a topic and notice all the subscriber subscribing this topic
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





