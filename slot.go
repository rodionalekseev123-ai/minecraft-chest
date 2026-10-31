package main

import "fmt"

type Slot struct { // Структура Slot:
	Name  string `json:"name"` // пустая строка — ячейка пуста
	Count int    `json:"count"`
}

func (s Slot) IsEmpty() bool {

	var empty bool

	if s.Name == "" {
		empty = true
	} else {
		empty = false
	}
	return empty
}

// сколько ещё можно положить в эту ячейку (не больше MaxStack)
func (s Slot) FreeSpace() int {
	Max := MaxStack

	Max -= s.Count
	if Max < 0 {
		Max = 0
	}

	return Max
}

// в ячейке лежит именно этот предмет
func (s Slot) SameItem(name string) bool {

	var flag bool

	if s.Name == name {
		flag = true
	} else {
		flag = false
	}
	return flag
}

// очистить ячейку
func (s Slot) Clear() {

	if s.Name != "" && s.Count > 0 {
		s.Name = ""
		s.Count = 0
	}
}

// положить сколько влезет; вернуть сколько положили
func (s *Slot) Add(name string, n int) (int, error) {

	var rest int //сколько положили
	if n <= 0 {
		return 0, fmt.Errorf("invalid value")
	}

	if s.IsEmpty() == true {
		s.Name = name
		s.Count += n
		if s.Count >= MaxStack {
			s.Count = MaxStack
			rest = MaxStack
			return rest, nil
		} else {
			rest = n
			return rest, nil
		}
	} else if s.SameItem(name) == true {
		if s.Count == MaxStack {
			return 0, nil
		} else {
			rest = s.Count
			s.Count += n
			if s.Count >= MaxStack {
				s.Count = MaxStack
				rest = s.Count - rest
				return rest, nil
			} else {
				rest = n
				return n, nil
			}
		}
	} else {
		return 0, nil
	}

}

func (s *Slot) Take(n int) (int, error) {

	cnt := s.Count

	s.Count -= n
	if s.Count > 0 {
		return n, nil
	} else if s.Count == 0 {
		s.Clear()
		return n, nil
	} else if s.Count < 0 {
		return cnt, nil
	} else {
		return 0, fmt.Errorf("bad_count")
	}
}

//
// Методы:
//   IsEmpty() bool
//   FreeSpace() int                 // сколько ещё можно положить в эту ячейку (не больше MaxStack)
//   SameItem(name string) bool      // в ячейке лежит именно этот предмет
//   Clear()                         // очистить ячейку
//   Add(name string, n int) (int, error)  // положить сколько влезет; вернуть сколько положили
//   Take(n int) (int, error)              // забрать; вернуть сколько забрали; если стало 0 — Clear
//
// Ошибки Add/Take при n <= 0: bad_count.
