package informations

type Service interface {
	CreateInformations(input CreateInformationsInput) (Informations, error)
	GetInformations(IDUser int) ([]Informations, error)
	GetInformationsByID(input GetInformationsDetailInput) (Informations, error)
	UpdateInformations(inputID GetInformationsDetailInput, input CreateInformationsInput) (Informations, error)
	DeleteInformation(inputID GetInformationsDetailInput) (Informations, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateInformations(input CreateInformationsInput) (Informations, error) {
	information := Informations{}
	information.IDUser = input.IDUser
	information.Npm = input.Npm
	information.Prodi = input.Prodi
	information.Ktp = input.Ktp
	information.Pekerjaan = input.Pekerjaan
	information.JenisPekerjaan = input.JenisPekerjaan
	information.Status = input.Status

	newInformation, err := s.repository.Save(information)

	if err != nil {
		return newInformation, err
	}

	return newInformation, nil
}

func (s *service) GetInformations(IDUser int) ([]Informations, error) {
	if IDUser != 0 {
		information, err := s.repository.FindByUserID(IDUser)
		if err != nil {
			return information, err
		}
		return information, nil
	}

	information, err := s.repository.FindAll()
	if err != nil {
		return information, err
	}
	return information, nil
}

func (s *service) GetInformationsByID(input GetInformationsDetailInput) (Informations, error) {
	information, err := s.repository.FindByID(input.ID)

	if err != nil {
		return information, err
	}

	return information, nil
}

func (s *service) UpdateInformations(inputID GetInformationsDetailInput, input CreateInformationsInput) (Informations, error) {
	information, err := s.repository.FindByID(inputID.ID)

	if err != nil {
		return information, err
	}

	// add this when you want to authorize user( client side )
	// if information.IDUser != input.User.ID {
	// 	return information, errors.New("Not Authorized As An Owner")
	// }

	information.IDUser = input.IDUser
	information.Npm = input.Npm
	information.Prodi = input.Prodi
	information.Ktp = input.Ktp
	information.Pekerjaan = input.Pekerjaan
	information.JenisPekerjaan = input.JenisPekerjaan
	information.Status = input.Status

	updatedInformation, err := s.repository.Update(information)

	if err != nil {
		return updatedInformation, err
	}

	return updatedInformation, nil
}

func (s *service) DeleteInformation(inputID GetInformationsDetailInput) (Informations, error) {

	deleteInformation, err := s.repository.Destroy(inputID.ID)

	if err != nil {
		return deleteInformation, err
	}

	return deleteInformation, nil

}
