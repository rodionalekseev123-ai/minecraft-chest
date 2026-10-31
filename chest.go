package main

import (
	"fmt"
)

// Структура Chest:
//
//	Slots []Slot `json:"slots"`  // ровно SlotCount ячеек, индексы 0 … SlotCount-1
type Chest struct {
	Slots []Slot `json:"slots"`
}

type SlotView struct {
	Index int
	Name  string
	Count int
}

func (c Chest) UsedSlots() int {

	cntUsed := 0
	items := c.Slots

	for i := 0; i < len(items); i++ {

		if items[i].Count > 0 {
			cntUsed++
		}
	}
	return cntUsed
}

func (c Chest) FreeSlots() int {

	cntFree := 0
	items := c.Slots

	for i := 0; i < len(items); i++ {

		if items[i].Count == 0 {
			cntFree++
		}
	}
	return cntFree

}

func (c Chest) TotalItems() int {
	total := 0
	items := c.Slots

	for i := 0; i < len(items); i++ {

		if items[i].Count > 0 {
			total += items[i].Count
		}
	}
	return total

}

func (c Chest) CountItem(name string) int {
	capacity := 0
	ourItems := c.Find(name)
	if len(ourItems) > 0 {
		for i := 0; i < len(ourItems); i++ {
			capacity += ourItems[i].Count
		}
	}
	return capacity
}

// занятые этим предметом ячейки, по возрастанию индекса
func (c Chest) Find(name string) []SlotView {

	var sameItems []SlotView

	for i := 0; i < len(c.Slots); i++ {
		if c.Slots[i].Name == name {
			var cntItem SlotView

			cntItem.Index = i
			cntItem.Name = c.Slots[i].Name
			cntItem.Count = c.Slots[i].Count

			sameItems = append(sameItems, cntItem)

		}
	}
	return sameItems
}

func (c Chest) Put(name string, count int) error {

	//Add() положить сколько влезет; вернуть сколько положили в слоте
	addCnt := count

	//посчитаем количество свободных ячеек
	freeSlost := c.FreeSlots()
	capacity := freeSlost * 64
	ourItems := c.Find(name)
	if len(ourItems) > 0 {
		for i := 0; i < len(ourItems); i++ {
			capacity += MaxStack - ourItems[i].Count
		}
	}

	if count <= 0 { //Валидация на введение число
		return fmt.Errorf("bad_count")
	} else if IsKnownItem(name) == false { //Проверяем есть предмет в списке items
		return fmt.Errorf("Item does not exist")
	} else if capacity < count {
		return fmt.Errorf("no_space")
	}
	if addCnt > 0 && len(c.Find(name)) > 0 { //проверяем есть ли ячейки, где уже лежит наш предмет
		ourItem := c.Find(name)

		for i := 0; i < len(ourItem); i++ {
			if ourItem[i].Count < MaxStack {
				idx := ourItem[i].Index
				added, err := c.Slots[idx].Add(name, addCnt)
				if err != nil {
					return err
				}
				addCnt -= added
				if addCnt <= 0 {
					break
				}
			}
		}
	}
	if addCnt > 0 {
		for i := 0; i < len(c.Slots); i++ {
			added, err := c.Slots[i].Add(name, addCnt)
			if err != nil {
				return err
			}
			addCnt -= added
			if addCnt <= 0 {
				break
			}
		}
	}
	return nil
}

func (c Chest) Take(name string, count int) error {

	takeCnt := count

	//проверяем сколько всего предметов в сундуке, которые нам нужны
	capacity := c.CountItem(name)

	ourItems := c.Find(name)

	if count > capacity {
		return fmt.Errorf("not_enough")
	} else if count <= 0 { //Валидация на введение число
		return fmt.Errorf("bad_count")
	} else if IsKnownItem(name) == false { //Проверяем есть предмет в списке items
		return fmt.Errorf("Item does not exist")
	}
	if takeCnt > 0 {
		for i := 0; i < len(ourItems); i++ {
			idx := ourItems[i].Index
			taked, err := c.Slots[idx].Take(takeCnt)
			takeCnt -= taked
			if err != nil {
				return err
			}
			if takeCnt <= 0 {
				break
			}
		}
	}
	return nil

}

// Методы:
//   UsedSlots() int
//   FreeSlots() int
//   TotalItems() int
//   CountItem(name string) int
//   Find(name string) []SlotView   // занятые этим предметом ячейки, по возрастанию индекса
//   Put(name string, count int) error
//   Take(name string, count int) error
//
// Put / Take меняют сундук.
//
// Put:
//   1) предмет должен быть из списка (иначе unknown_item)
//   2) count > 0 (иначе bad_count)
//   3) если свободного места под этот предмет не хватает на весь count — no_space, сундук не менять
//   4) сначала докладывать в уже существующие стаки этого предмета (с меньшего индекса)
//   5) остаток класть в пустые ячейки (с меньшего индекса), не больше MaxStack в каждую
//
// Take:
//   1) unknown_item / bad_count — как у Put
//   2) если предметов name в сундуке меньше count — not_enough, сундук не менять
//   3) забирать начиная с меньшего индекса; опустевшую ячейку очищать через Clear
//
// Структура SlotView (для Find):
//   Index int
//   Name  string
//   Count int
