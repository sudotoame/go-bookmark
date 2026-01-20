package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const menuText = `
- 1. Посмотреть закладки
- 2. Добавить закладку
- 3. Удалить закладку
- 4. Выход
`

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var isContinue bool
	m := map[string]string{}
	for !isContinue {
		fmt.Print(menuText)
		getInput, err := getUserInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch getInput {
		case 1:
			for key, value := range m {
				fmt.Println(key, value)
			}
		case 2:
			err := updateMap(m)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case 3:
			err := deleteMap(m)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case 4:
			isContinue = true
		}
	}
}

func deleteMap(m map[string]string) error {
	fmt.Print("Введите название для удаления: ")
	del := readLine()
	if del == "" {
		return errors.New("Название не может быть пустым!")
	}
	delete(m, del)
	fmt.Printf("Поле %q успешно удалено!", del)
	return nil
}

func updateMap(m map[string]string) error {
	fmt.Print("Введите название: ")
	key := readLine()
	fmt.Print("Введите адресс: ")
	url := readLine()

	if key == "" || url == "" {
		return errors.New("Название и адрес не могут быть пустыми!")
	}

	m[key] = url
	fmt.Printf("Поле %q успешно добавлено!", key)

	return nil
}

func getUserInput() (int, error) {
	var ch int
	_, err := fmt.Scan(&ch)
	if err != nil {
		return 0, err
	}
	if ch != 1 && ch != 2 && ch != 3 && ch != 4 {
		fmt.Println("Ошибка ввода! Введите число от 1 до 4")
		return 0, err
	}

	return ch, nil
}

func readLine() string {
	scanner.Scan()
	return scanner.Text()
}
