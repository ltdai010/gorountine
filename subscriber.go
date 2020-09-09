package main

type Subscriber struct {
	name string
	topicList []string
	receiver chan string
}

func (this *Subscriber) init()  {
	this.receiver = make(chan string)
}

func (this *Subscriber) subscribe(topic *Topic) {
	this.topicList = append(this.topicList, topic.name)
	topic.subscriber = append(topic.subscriber, this.name)
}


