package main

type Publisher struct {
	name string

}

func (this *Publisher) publish(topic *Topic, content string)  {
	topic.content = append(topic.content, content)
	topic.broadcast <- content
}
