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

func (this *Topic) notifyAll(subscriberList map[string]*Subscriber) bool {
	select {
		case s := <- this.broadcast:
			for _, i := range this.subscriber {
				subscriberList[i].receiver <- s
			}
		case <-quit:
			return true
		default:
			return false
	}
	return false
}