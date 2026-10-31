package main

import "fmt"

// 1. LoadChest("chest.json")
// 2. Напечатать UsedSlots и FreeSlots
// 3. Бесконечный цикл: читать команду, выполнять, пока не exit
// 4. После exit: UsedSlots, FreeSlots, TotalItems
// 5. SaveChest("chest.json", chest)
// 6. Напечатать содержимое chest.json одной строкой
// 7. LoadChest снова — напечатать UsedSlots и FreeSlots
//
// count и find — через var st Storage = &chest
// Ошибка: fmt.Println("error", err.Error())
// Успех put/take: fmt.Println("ok")

func main() {
	repo := FileRepository{"chest.json"}

	chest, err := repo.LoadChest()
	if err != nil {
		fmt.Println("error", err)
		return
	}

	cntUsed := chest.UsedSlots()
	freeSlost := chest.FreeSlots()
	//totalItems := chest.TotalItems()

	fmt.Println(chest.Slots)
	fmt.Println("Количество занятых ячеек:", cntUsed)
	fmt.Println("Количество свободных ячеек:", freeSlost)
	//fmt.Println("Количество всех предметов:", totalItems)

	chest.Put()
	fmt.Println("Сундук после изменений")
	fmt.Println(chest.Slots)

}

// func main() {
// 	var repo UserRepository

// 	println("choose 1 or 2 to select repository (console/file)")
// 	var choice int
// 	fmt.Scan(&choice)

// 	switch choice {
// 	case 1:
// 		//из консоли
// 		repo = ConsoleRepository{}
// 	case 2:
// 		repo = FileRepository{"user.json"}
// 	}

// 	user, err := repo.GetUser()
// 	if err != nil {
// 		println("something went wrong")
// 		return
// 	}

// 	println("result: ")
// 	println(user.Name)
// 	println(user.Age)
// }
