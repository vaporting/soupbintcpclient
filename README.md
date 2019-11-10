# SoupBinTCPClient
A simple SoupBinTCP client

## Execution
* Run
```
non-build:
go run main.go
build:
./soupbintcpclient
```
* Stop
    * send system interrupt(ctrl-C)

## Command line introduction
* server_addr  // set server ip address, default is 127.0.0.1
* server_port  // set server port, default is 30010
* client_port  // set client port, default is empty
```
e.g.
go run main.go -server_addr=127.0.0.1 -server_port=30010 -client_port=11111
```

## How to build
1. Clone this project
2. go to project folder
3. Run build command
    ```
    go build -o soupbintcp_client soupbintcpclient
    ```
4. get binary file: soupbintcp_client

## How to run tests
1. Clone this project
2. Softlink this project in go/src
    ```
    $ ln -s [project_path] [GOPATH]/src/soupbintcpclient
    ```
3. Go to project folder
4. Run go get command to get all dependencies
    ```
    $ go get -t -v ./...
    ```
5. Run test command
    ```
    $ go test -v ./...
    ```
   
## Check list
- [x] The client should be compiled to a binary that can run in alpine linux.
- [x] The binary can use arguments (or flags) to specify which server will be connected 
- [x] The client should exit when it receives an OS interrupt signal


- [x] Implement error handling logic and don't just panic
- [x] Write some unit tests
- [x] Do the following actions (in order) after the client connects to the server
    1. Send a client heartbeat at time 0ms
    2. Send a debug packet at time 1300ms
    3. Send client heartbeats if need
    4. Send an unsequenced data packet at time 4600ms and 5400ms 
    5. Close connection at time 10000ms

## Notice
### Deploy on alpine linux docker
if you want to deploy binary of client to alpine linux docker, you have to run below command first.
```
$ mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
```

* referece
    * [Installed Go binary not found in path on Alpine Linux Docker](https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker)
