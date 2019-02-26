package main

import "fmt"

type Item struct {
	title string
	body  string
}

var database []Item

func GetItemByName(title string) Item {
	var returnedItem Item

	for _, val := range database {
		if val.title == title {
			returnedItem = val
		}
	}

	return returnedItem
}

func AddItem(item Item) Item {
	database = append(database, item)
	return item
}

func EditItem(title string, edit Item) Item {
	var changedItem Item

	for idx, val := range database {
		if val.title == title {
			database[idx] = edit
			changedItem = edit
		}
	}

	return changedItem
}

func DeleteItem(item Item) Item {
	var deletedItem Item

	for idx, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:idx], database[idx+1:]...)
			deletedItem = item
			break
		}
	}

	return deletedItem
}

func main() {
	fmt.Println("initial database: ", database)
	a := Item{"first", "first test item"}
	b := Item{"second", "second test item"}
	c := Item{"third", "third test item"}

	AddItem(a)
	AddItem(b)
	AddItem(c)
	fmt.Println("second database: ", database)

	DeleteItem(b)
	fmt.Println("third database: ", database)

	EditItem("third", Item{"forth", "forth test item"})
	fmt.Println("forth database: ", database)

	x := GetItemByName("forth")
	y := GetItemByName("first")
	fmt.Println(x, y)

}
