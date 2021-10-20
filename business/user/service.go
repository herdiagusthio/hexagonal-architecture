package user

type service struct {
	repository Repository
}

//NewService Construct user service object
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) FindUserByID(id int) (*FindUser, error) {
	return s.repository.FindUserByID(id)
}
