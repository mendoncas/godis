package main

import (
	"log"
	"net"
	"github.com/mendoncas/godis/components"
)

func main() {
  // cache := make(map[string]string)
  cache := components.CreateTable(5)
  cache.Print()
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
    go connectionListenLoop(conn, cache)
  }
}

func connectionListenLoop(conn net.Conn, cache *components.Table) {
  log.Print("new connection from ", conn.RemoteAddr())
  for {
    received := make([]byte, 1024)
    _, err := conn.Read(received)
    if err != nil {
      log.Print("Closing connection from ", conn.RemoteAddr(),  err)
      conn.Close()
      return
    }
    handleMessage(string(received), conn, cache)
  }
}

func handleMessage(message string, conn net.Conn, cache *components.Table) {
    packet := components.Packet{}
    packet.FromString(string(message))
    switch packet.Command {
      case "PNG":
        conn.Write([]byte("pong")) 
      case "ADD":
        cache.Add(packet.Key, packet.Value) 
        // (*cache)[packet.Key] = packet.Value
      case "GET":
        node, _ := cache.Get(packet.Key)
        conn.Write([]byte(node.Value))
        // conn.Write([]byte((*cache)[packet.Key])) 
      case "DEL":
        cache.Remove(packet.Key)
        // delete(*cache, packet.Key)
      default:
        log.Print("Unknown command: ", packet.Command)
    }
    cache.Print()
}
