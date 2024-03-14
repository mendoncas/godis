package main

import (
	"log"
  "time"
	"net"
)

func main(){
  conn, err := net.Dial("tcp", "localhost:3000")
  defer conn.Close()
  go pingServer(conn)
  for{
    if err != nil {
      log.Print("Failed to connect to server!", err)
    }

    received := make([]byte, 1024)
    _, err = conn.Read(received)
    if err != nil {
      log.Print("Failed to read from connection!", err)
    }
    log.Print("Received: ", string(received))
  }
}

func pingServer(conn net.Conn){
  for {
    // conn.Write([]byte("PNG"))
    time.Sleep(5 * time.Second)
    conn.Write([]byte("ADDTESTE"))
  }
}
