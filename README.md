## 版本更新
---

### v0.0.1 
2014-05-08 viney
- 使用socket.io实现简易聊天室.
- 服务端使用go-socket.io实现.
- 客户端使用socket.io.js实现.
- 由于服务端没找到获取IP的方法，有修改go-socket.io一点代码

-  server.go文件
```go
    type SocketIOServer struct {
        ...
        *http.Request
    }
```
-  server.go文件
```go
    func (srv *SocketIOServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        srv.Request = r
        ...
    }
```
