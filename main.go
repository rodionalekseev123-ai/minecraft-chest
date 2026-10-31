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

	// 1. LoadChest("chest.json")
	chest, err := repo.LoadChest()
	if err != nil {
		fmt.Println("error", err)
		return
	}

	var st Storage = &chest

	//2. Напечатать UsedSlots и FreeSlots
	cntUsedBefore := st.UsedSlots()
	freeSlostBefore := st.FreeSlots()

	fmt.Println("Количество занятых ячеек:", cntUsedBefore)
	fmt.Println("Количество свободных ячеек:", freeSlostBefore)

	//3. Бесконечный цикл: читать команду, выполнять, пока не exit
	for {
		var chose string
		fmt.Println("Что вы хотите сделать?")
		fmt.Scan(&chose)
		switch chose {
		case "put":
			fmt.Println("Какой предмет вы хотите положить?")
			itemName := fmt.Scan(&itemName)
			fmt.Println("Сколько единиц?")
			itemCount := fmt.Scan(&itemCount)
			chest.Put(itemName, itemCount)

		}

	}

	// 4. После exit: UsedSlots, FreeSlots, TotalItems
	cntUsedAfter := st.UsedSlots()
	freeSlostAfter := st.FreeSlots()
	totalItems := chest.TotalItems()
	fmt.Println("Количество занятых ячеек:", cntUsedAfter)
	fmt.Println("Количество свободных ячеек:", freeSlostAfter)
	fmt.Println("Количество всех предметов:", totalItems)

	// 5. SaveChest("chest.json", chest)

	// 6. Напечатать содержимое chest.json одной строкой
	fmt.Println("Сундук после изменений")
	fmt.Println(chest.Slots)

	// 7. LoadChest снова — напечатать UsedSlots и FreeSlots

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
