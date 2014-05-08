package main

import (
	"errors"
	"sync"
	"time"

	"code.google.com/p/go-uuid/uuid"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type Account interface {
	Add(name string)
	Update(string, *User) error
	Del(uuid string) error
	FindAll() map[string]*User
}

type account struct {
	mux sync.Mutex
	// 用户信息列表
	list map[string]*User
}

type User struct {
	UUID       string
	Name       string
	CreateTime time.Time
}

func New() Account {
	return &account{
		mux:  sync.Mutex{},
		list: make(map[string]*User),
	}
}

// -------------------------
// Add

func (a *account) Add(name string) {
	// uuid
	uuid := uuid.New()
	user := &User{
		UUID:       uuid,
		Name:       name,
		CreateTime: time.Now(),
	}

	a.mux.Lock()
	a.list[uuid] = user
	defer a.mux.Unlock()
}

// ------------------------
// Update

func (a *account) Update(uuid string, user *User) error {
	a.mux.Lock()

	// 验证用户是否存在
	if _, ok := a.list[uuid]; !ok {
		defer a.mux.Unlock()
		return ErrUserNotFound
	}

	// update
	a.list[uuid] = user
	defer a.mux.Unlock()

	return nil
}

// ------------------------
// Del

func (a *account) Del(uuid string) error {
	a.mux.Lock()

	// 验证用户是否存在
	if _, ok := a.list[uuid]; !ok {
		defer a.mux.Unlock()
		return ErrUserNotFound
	}

	// update
	delete(a.list, uuid)
	defer a.mux.Unlock()

	return nil
}

// ------------------------
// FindAll

func (a *account) FindAll() map[string]*User {
	return a.list
}
