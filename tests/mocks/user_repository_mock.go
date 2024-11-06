package mocks

import (
	"errors"

	"github.com/tanvirislam06/go-api-test/internal/user"
)

type MockUserRepository struct {
	MockUsers map[int]*user.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		MockUsers: map[int]*user.User{
			1: {ID: 1, Name: "AliceInChains", Email: "aliceinchains@example.com"},
			2: {ID: 2, Name: "SystemOfADown", Email: "systemofadown@example.com"},
		},
	}
}

func (m *MockUserRepository) GetUserById(id int) (*user.User, error) {
	user, exists := m.MockUsers[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
