package entity

type OrderRepositorytInterface interface { // to all orders
	Save(order *Order) error
	GetTotal() (int, error)
}
