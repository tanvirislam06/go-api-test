package user

import "errors"

type User struct {
	ID    int
	Name  string
	Email string
}

type Repository struct {
	users map[int]*User
}

func NewRepository() *Repository {
	return &Repository{
		users: map[int]*User{
			1: {ID: 1, Name: "John", Email: "john@example.com"},
			2: {ID: 2, Name: "Jane", Email: "jane@example.com"},
		},
	}
}

func (r *Repository) GetUserById(id int) (*User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
