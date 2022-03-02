/*
   Implementation of Set API
*/

package goset

type Map map[string]struct{}
type Set struct{ m Map }

func CreateNewSet(size int) *Set {
    return &Set{make(Map, size)}
}

func Create(values ...string) *Set {
    s := CreateNewSet(len(values))
    for _, value := range values {
        s.Add(value)
    }
    return s
}

func (s *Set) Add(value string) {
    if !s.Contains(value) {
        s.m[value] = struct{}{}
    }
}

func (s *Set) Contains(value string) bool {
    if _, ok := s.m[value]; !ok {
        return false
    }
    return true
}

func (s *Set) Remove(value string) {
    if !s.Contains(value) {
        delete(s.m, value)
    }
}

func (s *Set) Size() int {
    return len(s.m)
}
