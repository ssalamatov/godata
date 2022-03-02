package gobtree

import (
	"testing"
)

func TestCreateTree(t *testing.T) {
    tree := Create()
    if t == nil {
        t.Errorf("Got nil pointer on new tree %p", tree)
    }

}
func TestInsertNode(t *testing.T) {
    /*
      Create Tree from 1 node
    */
    tree := Create()
    tree.Insert(10)

    isfound := tree.Exists(10)
    if !isfound {
      t.Errorf("Not found 10 node %v", isfound)
    }
  }

  func TestFindAbsentNode(t *testing.T) {
    /*
      Find absent value
    */
    tree := Create()
    tree.Insert(10)

    isfound := tree.Exists(5)
    if isfound {
      t.Errorf("Found unknown value %v", isfound)
    }
  }
