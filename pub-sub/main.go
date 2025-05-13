package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	Topic   string
	Content string
}

type Subscriber struct {
	ID   string
	Ch   chan Message
	Done chan struct{}
}

type Broker struct {
	mu     sync.RWMutex
	topics map[string][]*Subscriber
}

func NewBroker() *Broker {
	return &Broker{
		topics: make(map[string][]*Subscriber),
	}
}

func (b *Broker) Subscribe(topic string, sub *Subscriber) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.topics[topic] = append(b.topics[topic], sub)
	fmt.Printf("[Broker] Subscriber %s subscribed to %s\n", sub.ID, topic)
}

func (b *Broker) Unsubscribe(topic string, sub *Subscriber) {
	b.mu.Lock()
	defer b.mu.Unlock()
	subs := b.topics[topic]
	for i, s := range subs {
		if s.ID == sub.ID {
			b.topics[topic] = append(subs[:i], subs[i+1:]...)
			close(s.Ch)
			close(s.Done)
			fmt.Printf("[Broker] Subscriber %s unsubscribed from %s\n", sub.ID, topic)
			break
		}
	}
}

func (b *Broker) Publish(topic string, content string) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, sub := range b.topics[topic] {
		select {
		case sub.Ch <- Message{Topic: topic, Content: content}:
		default:
			fmt.Printf("[Broker] Subscriber %s's channel is full. Dropping message.\n", sub.ID)
		}
	}
}

func NewSubscriber(id string) *Subscriber {
	return &Subscriber{
		ID:   id,
		Ch:   make(chan Message, 10), // buffer messages
		Done: make(chan struct{}),
	}
}

func main() {
	broker := NewBroker()

	sub1 := NewSubscriber("A")
	sub2 := NewSubscriber("B")

	broker.Subscribe("news", sub1)
	broker.Subscribe("news", sub2)

	go func(sub *Subscriber) {
		for {
			select {
			case msg := <-sub.Ch:
				fmt.Printf("[Sub %s] Got message: %s\n", sub.ID, msg.Content)
			case <-sub.Done:
				return
			}
		}
	}(sub1)

	go func(sub *Subscriber) {
		for {
			select {
			case msg := <-sub.Ch:
				fmt.Printf("[Sub %s] Got message: %s\n", sub.ID, msg.Content)
			case <-sub.Done:
				return
			}
		}
	}(sub2)

	// Publish some messages
	broker.Publish("news", "Breaking: Go is awesome!")
	broker.Publish("news", "Latest update on Pub-Sub design.")

	time.Sleep(2 * time.Second)

	broker.Unsubscribe("news", sub1)

	broker.Publish("news", "Sub1 should not get this.")

	time.Sleep(2 * time.Second)
}
