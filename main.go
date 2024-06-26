package main

import (
	"fmt"
)

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	slice := make([]Message, 0)
	for _, m := range messages {
		if m.Type() == filterType {
			slice = append(slice, m)
		}
	}

	return slice
}

func main() {

	fmt.Println("hello world")
}
