package main

import (
	"net/http"
	"time"

	"engine/errors"
	"engine/log"
	socketio "github.com/googollee/go-socket.io"
)

var (
	// 实例化聊天室
	c = New()

	// 用户列表
	users = make(map[string]*socketio.NameSpace)

	ips = make(map[string]string)
)

// socket
var (
	sio = socketio.NewSocketIOServer(new(socketio.Config))
)

func connect(ns *socketio.NameSpace) {
	// 获取所有数据
	messages := c.FindAll()
	if err := ns.Emit("connect", messages); err != nil {
		log.Warn(errors.As(err))
		return
	}

	id := ns.Id()
	ip := Ip(sio.Request)
	content := "我上线了,大家在聊什么呢?"
	message := &Message{ip, Format(time.Now()), content}

	// go func(map[string]*socketio.NameSpace, *Message) {
	for k, ons := range users {
		if k != id {
			if err := ons.Emit("join", message); err != nil {
				log.Warn(errors.As(err, k))
				return
			}
		}
	}
	// }(users, message)

	if _, ok := users[id]; !ok {
		// 添加到列表
		users[id] = ns
		ips[id] = ip
	}
}

func disconnect(ns *socketio.NameSpace) {
	id := ns.Id()
	ip := Ip(sio.Request)
	content := "我已经下线,大家聊的开心!"
	message := &Message{ip, Format(time.Now()), content}

	for k, ons := range users {
		if k != id {
			if err := ons.Emit("quit", message); err != nil {
				log.Warn(errors.As(err, k))
				return
			}
		}
	}

	if _, ok := users[id]; ok {
		// 从列表删除
		delete(users, id)
		delete(ips, ip)
	}
}

func say(ns *socketio.NameSpace, content string) {
	ip := ips[ns.Id()]
	message := c.Add(ip, content)

	// go func(map[string]*socketio.NameSpace, *Message) {
	for k, ons := range users {
		if err := ons.Emit("say", message); err != nil {
			log.Warn(errors.As(err, k))
			return
		}
	}
	// }(users, message)
}

func main() {
	// socket
	sio.On("connect", connect)
	sio.On("disconnect", disconnect)
	sio.On("say", say)

	sio.Handle("/", http.FileServer(http.Dir("./public/")))

	// listen
	if err := http.ListenAndServe(":8080", sio); err != nil {
		panic(errors.As(err))
	}
}
