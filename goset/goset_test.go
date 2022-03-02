package goset

import (
    "fmt"
    "testing"
)

func KeyExists(s *Set, keys ...string) bool {
    for _, key := range keys {
        if _, ok := s.m[key]; !ok {
            return false
        }
    }
    return true
}

func Error(t *testing.T, msg string, value interface{}) {
    text := fmt.Sprintf(msg, value)
    t.Error(text)
}

func TestCreateSet(t *testing.T) {
    /*
        Set is properly initialized
    */
    s := Create("aaa", "bbb")
    if s == nil {
        Error(t, "Got nil pointer on new set: %p", s)
    }

    s = Create()
    if s == nil {
        Error(t, "Got nil pointer on new set: %p", s)
    }
}

func TestKeysValidSet(t *testing.T) {
    /*
        Only stored keys are presented in set
    */
    s := Create("aaa", "bbb")

    if !KeyExists(s, "aaa", "bbb") {
        t.Errorf("Missing keys in set %s", s)
    }

    if KeyExists(s, "", "1") {
        t.Errorf("Wrong keys in set %s", s)
    }
}

func TestEmptyKeySet(t *testing.T) {
    /*
        Empty string as a key
    */
    s := Create("")
    if !KeyExists(s, "") {
        t.Errorf("Missing keys in set %s", s)
    }
}

func TestEmptySizeSet(t *testing.T) {
    /*
        Set has length equal to the number of the stored keys
    */
    s := Create()

    if len(s.m) != 0 {
        t.Errorf("Invalid size of set %s", s)
    }
}

func TestUniqueItemsSet(t *testing.T) {
    /*
        Set does not contain duplicates of keys
    */
    s := Create("aaa", "aaa")

    if len(s.m) != 1 {
        t.Errorf("Invalid size of set %s", s)
    }
}

func TestContainsSet(t *testing.T) {
    s := Create("aaa", "ccc")

    if !s.Contains("aaa") {
        t.Errorf("Key is missing %s", s)
    }

    if !s.Contains("ccc") {
        t.Errorf("Key is missing %s", s)
    }

    if s.Contains("bbb") {
        t.Errorf("Wrong key %s", s)
    }

}

func TestAddToSet(t *testing.T) {
    s := Create()
    s.Add("aaa")
    s.Add("ccc")

    if !KeyExists(s, "aaa", "ccc") {
        t.Errorf("Keys is missing %s", s)
    }
    if KeyExists(s, "zzz") {
        t.Errorf("Wrong key %s", s)
    }

}
