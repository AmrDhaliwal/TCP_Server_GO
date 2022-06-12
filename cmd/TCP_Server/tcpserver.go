package tcpserver

import (
	"fmt"
	"net"
	"os"
	//"log" -Replace print statements with log functions
)

const (
	HOST = "localhost"
	PORT = "3333"
	TCP  = "tcp" 
)

func StartServer() {
    // Listen for incoming connections.
    l, err := net.Listen(TCP, HOST+":"+PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + HOST + ":" + PORT)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleConnections(conn)
    }
}

func handleConnections(conn net.Conn) {
	fmt.Println("Entered handleConnections func....")
	fmt.Println(conn)
}



