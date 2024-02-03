package main

import (
	"fmt"
)

type Node struct {
	value int32
	next  *Node
}

func main() {
	var head *Node
	for {
		fmt.Println("1 - Добавить ноду")
		fmt.Println("2 - Вывести список")
		var res string
		_, err := fmt.Scan(&res)
		if err != nil {
			return
		}
		switch res {
		case "1":
			var val int32
			fmt.Println("Какое значение?")
			_, err := fmt.Scan(&val)
			if err != nil {
				return
			}
			if head == nil {
				head = &Node{value: val, next: nil}
			} else {
				current := head
				for current.next != nil {
					current = current.next
				}
				current.next = &Node{value: val, next: nil}
			}

		case "2":
			fmt.Println("Давай тогда покажу что ты натыкал")
			curr := head
			i := 1
			for {
				if curr != nil {
					fmt.Printf("Нода %#v со значением: %#v\n", i, curr.value)
					if curr.next != nil {
						curr = curr.next
						i++
					} else {
						break
					}
				}
			}
		default:
			fmt.Println("смотри че жмешь")

		}
	}
}
