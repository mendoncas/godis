package main

import (
	"log"
	"math/rand"
	"net"
	"time"
	"github.com/mendoncas/godis/components"
  "github.com/thanhpk/randstr"
)

func main(){
  conn, _ := getConnection()
  defer conn.Close()
  go pingServer(conn)
  for{
    received := make([]byte, 1024)
    _, err := conn.Read(received)
    if err != nil {
      log.Print("Failed to read from connection!", err)
      log.Print("Retrying connection...")
      conn, _ = getConnection()
      time.Sleep(5 * time.Second)
    }
    log.Print("Received: ", string(received))
  }
}

func getConnection() (net.Conn, error) {
  conn, err := net.Dial("tcp", "localhost:3000")
  if err != nil {
    log.Print("Failed to connect to server!", err)
  }
  return conn, err
}

func pingServer(conn net.Conn){
  commands := [3]string{"DEL", "ADD", "GET"}

  for {
    packet := components.Generate(commands[rand.Intn(3)], randstr.String(10), randstr.String(10))
    log.Print(randstr.String(10))
    conn.Write([]byte(packet.ToString()))
    time.Sleep(1 * time.Second)
  }
}
