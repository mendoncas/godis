package main

import (
	"log"
	"strconv"
)

type Packet struct {
  // command+payloadSize+payload
  // DEL0007PAYLOAD
  command string; // length == 3
  payloadSize int; 
  payload string; 
}

func (p *Packet) ToString() string {
  strPayloadSize := leftpad(strconv.Itoa(p.payloadSize), "0", 4)
  return  p.command +  strPayloadSize + p.payload
}

func (p *Packet) FromString(s string) {
  p.command = s[0:3]
  p.payloadSize, _ = strconv.Atoi(s[3:7])
  p.payload = s[7:p.payloadSize+7]
  log.Print("p.command: ", p.command, " p.payloadSize: ", p.payloadSize, " p.payload: ", p.payload)
}

func leftpad(s string, pad string, length int) string {
  for i := len(s); i < length; i++ {
    s = pad + s
  }
  return s
}

