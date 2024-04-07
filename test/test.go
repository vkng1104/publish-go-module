package main

import (
	"fmt"

	"github.com/vkng1104/datelist"
)

func main() {
	client := datelist.Client("WGQqyEjfXLtK74K5YuxrHn7v")

	calendars, _ := client.ListCalendars(nil)

	filters := make(map[string]string)
	filters["calendar_id"] = "441"
	filters["name"] = "Table"

	products, _ := client.ListProducts(filters)

	filters = make(map[string]string)
	filters["email"] = "test@test.com"
	filters["calendar_id"] = "441"
	filters["from"] = "2021-08-04T04:51:59.945Z"
	filters["to"] = "2021-08-30T04:51:59.945Z"
	slots, _ :=
		client.ListBookedSlots(filters)

	fmt.Println(calendars)
	fmt.Println(products)
	fmt.Println(slots)

	data := make(map[string]string)
	data["email"] = "test2@test.com"
	fmt.Println(client.UpdateBookedSlot((slots[0]["id"]).(float64), data))
	data = make(map[string]string)
	data["email"] = "test@test.com"
	fmt.Println(client.UpdateBookedSlot((slots[0]["id"]).(float64), data))
}
