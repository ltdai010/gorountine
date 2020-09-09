package main

type Topic struct {
	name string
	content []string
	broadcast chan string
	subscriber []string
}

func (this *Topic) init() {
	this.broadcast = make(chan string)
}

func (this *Topic) notifyAll(subscriberList map[string]*Subscriber, flag *bool) {
	select {
		case s := <- this.broadcast:
			for _, i := range subscriberList {
				i.receiver <- s
			}
		case <-quit:
			*flag = true
			return
	}
	*flag = false
	return
}