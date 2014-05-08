package main

import (
	"sync"
	"time"
)

type Chat interface {
	// 添加聊天内容
	Add(ip, content string) *Message
	FindAll() []*Message
}

type chat struct {
	mux      sync.Mutex
	messages []*Message
}

type Message struct {
	IP         string
	CreateTime string
	Content    string
}

func New() Chat {
	return &chat{
		mux:      sync.Mutex{},
		messages: []*Message{},
	}
}

// -------------------------
// Add

func (c *chat) Add(ip, content string) *Message {
	c.mux.Lock()
	defer c.mux.Unlock()

	message := &Message{
		IP:         ip,
		CreateTime: Format(time.Now()),
		Content:    content,
	}
	c.messages = append(c.messages, message)

	return message
}

// -------------------------
// FindAll

func (c *chat) FindAll() []*Message {
	c.mux.Lock()
	defer c.mux.Unlock()

	return c.messages
}
