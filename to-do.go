package main

import (
	"fmt"
	"os"
)	
	var map_index int = 1
	var list_value string
	var list = make(map[int]string)

func main() {
for{
	choose()
	}
}

func addTask()  {
	fmt.Printf("введите %d задания:", map_index)
	fmt.Scanln(&list_value)
	list[map_index] = list_value
	map_index++
}

func deleteTask()  {
	fmt.Printf("выберите номер задания чтобы удалить:")
	var b int 
	fmt.Scanf("%d\n", &b)

	delete(list, b)
}

func returnTask(){
	for id, task := range list {
		fmt.Printf("%d. %s\n", id, task)
	}
}

func off(){
	os.Exit(0)
}

func choose()   {
	fmt.Println("1:Добавать задания.")
	fmt.Println("2:Удалить задания.")
	fmt.Println("3:Вывод резельтата.")
	fmt.Println("4:Остановить ВСЕ!")

	fmt.Printf("выберите действие:")
	var i int 
	fmt.Scanf("%v\n", &i) 

	switch i {
		case 1:
			addTask()
		case 2:
			deleteTask()
		case 3:
			returnTask()
		case 4:
			off()
	}
}