package main

import "fmt"

func main () {

	names := []string{"Anna", "Boris", "Valentin"}

	var letter string
	fmt.Println("Введите букву: ")
	fmt.Scanf("%s\n", &letter)

	index := int(letter[0] - 'A') // преобразование буквы в индекс
	if len(letter) != 1 {
		fmt.Println("Вы ввели больше одного символа!")
		return
		} 

	if index >= 0 && index < len(names) {
		fmt.Printf("%s", names[index])
	} else {
		fmt.Println("С такой буквы нет имени в базе")
	}
}