package agregator

type User struct {
	ID   int
	Name string
}

type UserService struct {
	// ...
}

func (s *UserService) GetUser(id int) (*User, error) {
	// Aquí se implementa la lógica para obtener los datos del usuario del ID especificado.
	return nil, nil
}
