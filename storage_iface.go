package main

type Storage interface {
	CountItem(name string) int
	UsedSlots() int
	FreeSlots() int
	Find(name string) []SlotView
}

// Интерфейс Storage:
//   CountItem(name string) int
//   UsedSlots() int
//   FreeSlots() int
//   Find(name string) []SlotView
//
// Эти же методы должны быть у Chest — тогда переменной типа Storage
// можно присвоить &chest.
