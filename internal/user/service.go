package user

type Service struct {
	UserRepository *Repository
}

func NewUserService(userRepository *Repository) *Service {
	return &Service{
		UserRepository: userRepository,
	}
}

func (s *Service) GetUserById(id int) (*User, error) {
	user, error := s.UserRepository.GetUserById(id)
	return user, error
}
