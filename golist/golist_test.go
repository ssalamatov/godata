package golist

import (
  "testing"
)

func TestCreateList(t *testing.T) {
  /*
    List is properly initialized
  */
  l := Create()

  if l == nil {
    t.Errorf("Got nil pointer on new list %p", l)
  }
}

func TestInsertNode(t *testing.T) {
  /*
    Create List from 1 node
    Head and tail contain value from that node
  */
  l := Create()
  l.Insert("aaa")

  if l.Head().value != "aaa" {
    t.Errorf("Invalid head %v", l.head)
  }

  if l.Tail().value != "aaa" {
    t.Errorf("Invalid tail %v", l.head)
  }
}

func TestFindNode(t *testing.T) {
  /*
    Find node by value
  */
  l := Create()
  l.Insert("aaa")

  l.Insert("bbb")
  testnode := l.Head()

  l.Insert("ccc")

  found := l.Find("bbb")
  if testnode != found {
    t.Errorf("Found invalid node %v", found)
  }
}

func TestFindAbsentNode(t *testing.T) {
  /*
    Find absent value
  */
  l := Create()
  l.Insert("aaa")

  found := l.Find("bbb")
  if found != nil {
    t.Errorf("Got unknown node %v", found)
  }
}
