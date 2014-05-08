package main

import (
	"net/http"
	"strings"
	"time"
)

func Format(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// 返回请求IP地址
func Ip(req *http.Request) string {
	ip := req.Header.Get("X-Forwarded-For")
	if len(ip) == 0 {
		ip = strings.Split(req.RemoteAddr, ":")[0]
	}
	return ip
}
