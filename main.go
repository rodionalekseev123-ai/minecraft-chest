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
workLoop:
	for {
		var chose string

		fmt.Println("Что вы хотите сделать?")
		fmt.Scan(&chose)
		switch chose {
		case "put":
			var itemName string
			var itemCount int

			fmt.Println("Какой предмет вы хотите положить?")
			fmt.Scan(&itemName)
			fmt.Println("Сколько единиц?")
			fmt.Scan(&itemCount)
			err := chest.Put(itemName, itemCount)
			if err != nil {
				fmt.Println("error", err)
			} else {
				fmt.Println("ok")
			}
		case "take":
			var itemName string
			var itemCount int

			fmt.Println("Какой предмет вы хотите взять?")
			fmt.Scan(&itemName)
			fmt.Println("Сколько единиц?")
			fmt.Scan(&itemCount)
			err := chest.Take(itemName, itemCount)
			if err != nil {
				fmt.Println("error", err)
			} else {
				fmt.Println("ok")
			}
		case "count":
			var itemName string

			fmt.Println("Какой предмет ищите?")
			fmt.Scan(&itemName)
			itemCount := st.CountItem(itemName)
			fmt.Println(itemCount)

		case "find":
			var itemName string

			fmt.Println("Какой предмет ищите?")
			fmt.Scan(&itemName)
			fmt.Println(len(st.Find(itemName)))
			items := st.Find(itemName)
			fmt.Println(items)

		case "exit":
			break workLoop

		default:
			fmt.Println("error unknown_command")
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
	err = repo.SaveChest(chest)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	// 6. Напечатать содержимое chest.json одной строкой
	fmt.Println("Сундук после изменений")
	fmt.Println(chest.Slots)

	// 7. LoadChest снова — напечатать UsedSlots и FreeSlots
	chest, err = repo.LoadChest()
	if err != nil {
		fmt.Println("error", err)
		return
	}
	cntUsedEnd := st.UsedSlots()
	freeSlostEnd := st.FreeSlots()
	fmt.Println("Количество занятых ячеек:", cntUsedEnd)
	fmt.Println("Количество свободных ячеек:", freeSlostEnd)
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
