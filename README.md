# giniewebapp 
simple rest api build on golang using gin

# Install

## Clone
```shell script
$ git clone https://github.com/amjad489/giniewebapp
``` 

## Build
```shell script
$ go install
```

## Run
```shell script
$ giniewebapp
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2020/04/06 - 16:21:46 | 200 |      58.181µs |             ::1 | GET      "/ping"
[GIN] 2020/04/06 - 16:22:05 | 200 |       73.01µs |             ::1 | GET      "/ping"
```
## Test
```shell script
  $ curl -v 'http://localhost:8080/ping'
  *   Trying ::1...
  * TCP_NODELAY set
  * Connected to localhost (::1) port 8080 (#0)
  > GET /ping HTTP/1.1
  > Host: localhost:8080
  > User-Agent: curl/7.64.1
  > Accept: */*
  > 
  < HTTP/1.1 200 OK
  < Content-Type: application/json; charset=utf-8
  < Date: Mon, 06 Apr 2020 12:22:05 GMT
  < Content-Length: 18
  < 
  * Connection #0 to host localhost left intact
  {"message":"pong"}* Closing connection 0

```

