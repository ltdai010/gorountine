package main

import "fmt"

var (
	quit = make(chan int)
)

func main() {
	listSubscriber := make(map[string]*Subscriber)
	listTopic := make(map[string]*Topic)
	//create topic
	listTopic["test"] = &Topic{name: "test"}
	top := listTopic["test"]
	top.init()
	//crate subscriber
	listSubscriber["Dai"] = &Subscriber{name: "Dai"}
	sub := listSubscriber["Dai"]
	sub.init()
	sub.subscribe(top)
	//print notice of this subscriber
	go printNotice(sub)
	//publish a content
	go publish(top, "publish this content")
	//send to all subscriber
	for {
		if notifyAll(top, listSubscriber) {
			return
		}
	}
}


func notifyAll(topic *Topic, subscriberList map[string]*Subscriber) bool {
	select {
		case s := <- topic.broadcast:
			for _, i := range subscriberList {
				i.receiver <- s
			}
		case <-quit:
			return true
	}
	return false
}

func publish(topic *Topic, content string)  {
	topic.content = append(topic.content, content)
	topic.broadcast <- content
}

func printNotice(subscriber *Subscriber)  {
	fmt.Println(<-subscriber.receiver)
	quit <- 0
}





