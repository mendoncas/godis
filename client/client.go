package main

import (
	"log"
  "time"
	"net"
  "math/rand"
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
  commands := [6]string{
    "ADD000402OITESTE", "GET000002OI", "DEL000002OI", 
    "ADD000605FRUTABANANA", "GET000005FRUTA", "DEL000005FRUTA",}
  for {
    conn.Write([]byte(commands[rand.Intn(len(commands))]))
    time.Sleep(1 * time.Second)
  }
}
