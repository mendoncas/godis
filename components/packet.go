package components

import (
	"log"
	"strconv"
)

type Packet struct {
  // command+payloadSize+payload
  // DEL000701PAYLOAD
  payloadSize int; 
  keySize int;
  payload string; 
  
  Command string; // length == 3
  Key string;
  Value string;
}

func (p *Packet) ToString() string {
  strPayloadSize := leftpad(strconv.Itoa(p.payloadSize), "0", 4)
  return  p.Command +  strPayloadSize + p.payload
}

func Generate (command string, key string, value string) *Packet {
  payload := key + value
  return &Packet{
    Command: command,
    Key: key,
    Value: value,
    payload: payload,
    payloadSize: len(payload),
    keySize: len(key),
  }
}

func (p *Packet) FromString(s string) {
  p.Command = s[0:3]
  p.payloadSize, _ = strconv.Atoi(s[3:7])
  p.keySize, _ = strconv.Atoi(s[7:9])
  p.payload = s[9:p.payloadSize+9]
  p.Key = s[9:p.keySize+9]
  p.Value = s[p.keySize+9:]
  log.Println("Command: ", p.Command, p.Key, p.Value)
}

func leftpad(s string, pad string, length int) string {
  for i := len(s); i < length; i++ {
    s = pad + s
  }
  return s
}
