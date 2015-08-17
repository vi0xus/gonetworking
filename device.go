package main

import (
        "net"
        "fmt"
)

func main() {
    addr := net.UDPAddr{
        Port: 7001,
        IP: net.ParseIP("127.0.0.1"),
    }

    sock, _ := net.ListenUDP("udp", &addr)
    defer sock.Close()

    fmt.Printf("Listening...")

    var buf [512]byte

    for {
        n, addr, err := sock.ReadFromUDP(buf[0:])
        fmt.Println("Received ", string(buf[0:n]), " from ", addr)

        if err != nil {
            fmt.Println("Error: ", err)
        }
    }
}