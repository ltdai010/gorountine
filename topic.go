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
