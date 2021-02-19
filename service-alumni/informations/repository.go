package informations

import "gorm.io/gorm"

type Repository interface {
	Save(informations Informations) (Informations, error)
	FindAll() ([]Informations, error)
	FindByID(ID int) (Informations, error)
	FindByUserID(IDUser int) ([]Informations, error)
	Update(informations Informations) (Informations, error)
	Destroy(ID int) (Informations, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(informations Informations) (Informations, error) {
	err := r.db.Create(&informations).Error

	if err != nil {
		return informations, err
	}

	return informations, nil
}

func (r *repository) FindAll() ([]Informations, error) {
	var information []Informations
	err := r.db.Find(&information).Error

	if err != nil {
		return information, err
	}
	return information, nil
}

func (r *repository) FindByID(ID int) (Informations, error) {
	var information Informations
	err := r.db.Where("id = ?", ID).Find(&information).Error

	if err != nil {
		return information, err
	}
	return information, nil
}

func (r *repository) FindByUserID(IDUser int) ([]Informations, error) {
	var information []Informations
	err := r.db.Where("id_user = ?", IDUser).Find(&information).Error

	if err != nil {
		return information, err
	}
	return information, nil
}

func (r *repository) Update(informations Informations) (Informations, error) {
	err := r.db.Save(&informations).Error

	if err != nil {
		return informations, err
	}
	return informations, nil
}

func (r *repository) Destroy(ID int) (Informations, error) {
	var information Informations
	err := r.db.Where("id = ?", ID).Delete(&information).Error

	if err != nil {
		return information, err
	}

	return information, nil
}
