package agregator

type Order struct {
	ID     int
	UserID int
	Total  float64
}
type OrderService struct {
	// ...
}

func (s *OrderService) GetOrdersByUserID(userID int) ([]*Order, error) {
	// Aquí se implementa la lógica para obtener los pedidos del usuario con el ID especificado.
	return nil, nil
}
