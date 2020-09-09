package main

import (
	"fmt"
	"time"
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
	pub := listPublisher["Dia"]
	//publish a content
	go pub.publish(top, "pub 1 publish this content")
	//print notice of this subscriber
	go printNotice(listSubscriber)
	//send to all subscriber
	notifyAll(top, listSubscriber)
	time.Sleep(500*time.Millisecond)
}


func notifyAll(topic *Topic, subscriberList map[string]*Subscriber)  {
	select {
		case s := <- topic.broadcast:
			for _, i := range subscriberList {
				i.receiver <- s
			}
		case <-quit:
			return
	}
	return
}


func printNotice(subscriberList map[string]*Subscriber)  {
	for _, i := range subscriberList {
		fmt.Println(<-i.receiver + " to " + i.name)
	}
	quit <- 0
}





