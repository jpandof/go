package agregator

type UserOrders struct {
	User   *User
	Orders []*Order
}

type AggregatorService struct {
	UserService  *UserService
	OrderService *OrderService
}

// el patrón de agregador es una herramienta útil
// para simplificar las interacciones del cliente en una arquitectura de microservicios
// y para mejorar la eficiencia al evitar que los clientes tengan
// que hacer múltiples llamadas a la API y consolidar los datos por sí mismos.
func (s *AggregatorService) GetUserWithOrders(id int) (*UserOrders, error) {
	user, err := s.UserService.GetUser(id)
	if err != nil {
		return nil, err
	}

	orders, err := s.OrderService.GetOrdersByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	return &UserOrders{
		User:   user,
		Orders: orders,
	}, nil
}
