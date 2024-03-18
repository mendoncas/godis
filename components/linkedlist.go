package components

import (
	"errors"
	"log"
)

type LinkedList struct {
  FirstNode *ListNode
} 

func (list *LinkedList) Add(key string, value string) {
  if list.FirstNode == nil {
    list.FirstNode = &ListNode{Key: key, Value: value}
  } else{
    for node := list.FirstNode; node != nil; node = node.Next {
      if node.Next == nil {
        node.Next = &ListNode{Key: key, Value: value, Prev: node, Next: nil}
        log.Print("Added new node: ", node.Next)
        break
      } 
    }
  }
}

func (list *LinkedList) Get(key string) (*ListNode, error) {
  for node := list.FirstNode; node != nil; node = node.Next {
    if node.Key == key {
      return node, nil
    }
  }
  return &ListNode{}, errors.New("Key not found")
}

func (list *LinkedList) Remove(key string) error {
  node, err := list.Get(key)
  if err != nil {
    return err
  }
  if node == list.FirstNode {
    list.FirstNode = node.Next
  }
  if node.Prev != nil {
    node.Prev.Next = node.Next
  }
  if node.Next != nil {
    node.Next.Prev = node.Prev
  }
  return nil
}

func (list *LinkedList) Print() {
  str := ""
  for node := list.FirstNode; node != nil; node = node.Next {
    str += node.Key + ": " + node.Value + " -> "
  }
  log.Print(str)
}
