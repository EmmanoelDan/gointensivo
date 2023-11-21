package main

import "github.com/EmmanoelDan/gointensivo.git/internal/entity"

func main() {
	order, err := entity.NewOrder("1", -10, 1)

	if err != nil {
		println(err.Error())
	} else {
		println(order.ID)
	}
}
