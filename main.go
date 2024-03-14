package main

import (
	"log"
	"net"
	// "reflect"
)


func main() {
  cache := make(map[string]string)
  ln, err := net.Listen("tcp", ":3000")
  if err != nil {
    log.Panic("Failed to listen on port!", err)
  }
  for{
    conn, err := ln.Accept()
    defer conn.Close()
    if err != nil {
      log.Panic("Failed to accept connection!", err)
    }
    go connectionListenLoop(conn, &cache)
  }
}

func connectionListenLoop(conn net.Conn, cache *map[string]string) {
  log.Print("new connection from ", conn.RemoteAddr())
  for {
    received := make([]byte, 1024)
    _, err := conn.Read(received)
    if err != nil {
      log.Print("Closing connection from ", conn.RemoteAddr(),  err)
      conn.Close()
      return
    }
    log.Print("Received: ", string(received), " from ", conn.RemoteAddr())
    command := string(received[:3])
    switch command {
      case "PNG":
        conn.Write([]byte("pong")) 
      case "ADD":
        conn.Write([]byte("adding!")) 
      default:
        log.Print("Unknown command: ", command)
    }
    var st string = string(received[:])
    log.Print(len(st))
  }
}
