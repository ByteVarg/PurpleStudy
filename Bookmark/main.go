package main

import "fmt"

func viewBookmark(bookmarks map[string]string) {

	for key, value := range bookmarks {
		fmt.Printf("%s : %s", key, value)
	}

}

func addBookmark(bookmarks map[string]string) {

	var key, value string
	fmt.Println("Укажите название закладки: ")
	fmt.Scanln(&key)
	fmt.Println("Укажите ссылку: ")
	fmt.Scanln(&value)
	bookmarks[key] = value
}

func deleteBookmark(bookmarks map[string]string) {
	var key string
	fmt.Println("Укажите имя закладки, которую хотите удалить:")
	fmt.Scanln(&key)
	if _, ok := bookmarks[key]; ok {
		delete(bookmarks, key)

	} else {
		fmt.Println("Такой закладки не существует")

	}
}
func main() {

	bookmarks := map[string]string{"TestName": "TestLink"}
	var userInput int
	fmt.Println(`Добро пожаловать в приложение "Закладки"!`)
	fmt.Println(`
	****************************
	* Меню:                    *
	* 1. Просмотреть закладку  *
	* 2. Добавить закладку     *
	* 3. Удалить закладку      *
	* 4. Выход                 *
	****************************
	`)
MainCycle:
	for {

		fmt.Print("\nВведите пункт меню: ")
		fmt.Scan(&userInput)

		switch userInput {
		case 1:
			viewBookmark(bookmarks)
		case 2:
			addBookmark(bookmarks)
			fmt.Println("Закладка добавлена!")
		case 3:
			deleteBookmark(bookmarks)
			fmt.Println("Закладка удалена!")
		case 4:
			fmt.Println("Программа завершила свою работу.")
			break MainCycle
		default:
			fmt.Println("Неверный пункт меню. Попробуйте снова.")

		}
	}

}
