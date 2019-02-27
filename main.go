package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	title string
	body  string
}

type API int

var database []Item

func (a *API) GetItemByName(title string, reply *Item) error {
	var returnedItem Item

	for _, val := range database {
		if val.title == title {
			returnedItem = val
		}
	}

	*reply = returnedItem
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)

	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changedItem Item

	for idx, val := range database {
		if val.title == edit.title {
			database[idx] = edit
			changedItem = edit
		}
	}

	*reply = changedItem
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var deletedItem Item

	for idx, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:idx], database[idx+1:]...)
			deletedItem = item
			break
		}
	}

	*reply = deletedItem
	return nil
}

func main() {
	// fmt.Println("initial database: ", database)
	// a := Item{"first", "first test item"}
	// b := Item{"second", "second test item"}
	// c := Item{"third", "third test item"}

	// AddItem(a)
	// AddItem(b)
	// AddItem(c)
	// fmt.Println("second database: ", database)

	// DeleteItem(b)
	// fmt.Println("third database: ", database)

	// EditItem("third", Item{"forth", "forth test item"})
	// fmt.Println("forth database: ", database)

	// x := GetItemByName("forth")
	// y := GetItemByName("first")
	// fmt.Println(x, y)

	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}

	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}

}
