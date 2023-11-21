package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.validate()

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (order *Order) validate() error {
	if order.ID == "" {
		return errors.New("id is required")
	}

	if order.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	if order.Tax <= 0 {
		return errors.New("incorrect tax")
	}

	return nil
}

func (order *Order) CalculateFinalPrice() {
	order.FinalPrice = order.Price + order.Tax
}
